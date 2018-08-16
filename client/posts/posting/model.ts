import { message, send, handlers } from "../../connection"
import { Post } from "../model"
import { ImageData, PostData } from "../../common"
import FormView from "./view"
import { posts, storeMine, page, storeSeenPost, boardConfig } from "../../state"
import { postSM, postEvent, postState } from "."
import { extend } from "../../util"
import { SpliceResponse } from "../../client"
import { FileData } from "./upload"
import identity, { newAllocRequest } from "./identity"

// Form Model of an OP post
export default class FormModel extends Post {
	public inputBody = ""
	public view: FormView

	// Pass and ID, if you wish to hijack an existing model. To create a new
	// model pass zero.
	constructor() {
		// Initialize state
		super({
			id: 0,
			op: page.thread,
			editing: true,
			deleted: false,
			banned: false,
			sage: false,
			sticky: false,
			locked: false,
			time: Math.floor(Date.now() / 1000),
			body: "",
			name: "",
			auth: "",
			trip: "",
			state: {
				spoiler: false,
				quote: false,
				code: false,
				bold: false,
				italic: false,
				red: false,
				blue: false,
				haveSyncwatch: false,
				successive_newlines: 0,
				iDice: 0,
			},
		})
	}

	// Append a character to the model's body and reparse the line, if it's a
	// newline
	public append(code: number) {
		this.body += String.fromCodePoint(code)
	}

	// Remove the last character from the model's body
	public backspace() {
		this.body = this.body.slice(0, -1)
	}

	// Splice the last line of the body
	public splice(msg: SpliceResponse) {
		this.spliceText(msg)
	}

	// Compare new value to old and generate appropriate commands
	public parseInput(val: string): void {
		// Handle live update toggling
		if (postSM.state === postState.draft && !identity.live) {
			return
		}

		const old = this.inputBody

		// Rendering hack shenanigans - ignore
		if (old === val) {
			return
		}

		const lenDiff = val.length - old.length,
			exceeding = old.length + lenDiff - 2000

		// If exceeding max body length, shorten the value, trim input and try
		// again
		if (exceeding > 0) {
			this.view.trimInput(exceeding)
			return this.parseInput(val.slice(0, -exceeding))
		}

		// Remove any lines past 30
		const lines = val.split("\n")
		if (lines.length - 1 > 100) {
			const trimmed = lines.slice(0, 100).join("\n")
			this.view.trimInput(val.length - trimmed.length)
			return this.parseInput(trimmed)
		}

		if (postSM.state === postState.draft) {
			this.requestAlloc(val, null)
		} else if (lenDiff === 1 && val.slice(0, -1) === old) {
			// Commit a character appendage to the end of the line to the server
			const char = val.slice(-1);
			this.inputBody += char
			this.send(message.append, char.codePointAt(0))
		} else if (lenDiff === -1 && old.slice(0, -1) === val) {
			// Send a message about removing the last character of the line to
			// the server
			this.inputBody = this.inputBody.slice(0, -1)
			this.send(message.backspace, null)
		} else {
			this.commitSplice(val)
		}
	}

	// Optionally buffer all data, if currently disconnected
	private send(type: message, msg: any) {
		if (postSM.state !== postState.halted) {
			send(type, msg)
		}
	}

	// Commit any other input change that is not an append or backspace
	private commitSplice(v: string) {
		// Convert to arrays of chars to deal with multibyte unicode chars
		const old = [...this.inputBody],
			val = [...v],
			start = diffIndex(old, val),
			till = diffIndex(
				old.slice(start).reverse(),
				val.slice(start).reverse(),
			)

		this.send(message.splice, {
			start,
			len: old.length - till - start,
			// `|| undefined` ensures we never slice the string as [:0]
			text: val.slice(start, -till || undefined).join(""),
		})
		this.inputBody = v
	}

	// Close the form and revert to regular post. Cancel also erases all post
	// contents.
	public commitClose(cancel: boolean) {
		this.abandon()
		this.send(message.closePost, !!cancel) // Ensure boolean type
		if (cancel) {
			this.view.hide()
		}
	}

	// Turn post form into a regular post, because it has expired after a
	// period of posting ability loss
	public abandon() {
		this.view.cleanUp()
		this.closePost()
	}

	// Add a link to the target post in the input
	public addReference(id: number, sel: string) {
		let s = "";
		const old = this.view.input.value;
		const newLine = !old || old.endsWith("\n");

		if (sel) {
			if (!newLine) {
				s += "\n"
			}
		} else if (!newLine && old[old.length - 1] !== " ") {
			s += " "
		}
		s += `>>${id} `

		if (!sel) {
			// If starting from a new line, insert newline after post link
			if (newLine) {
				s += "\n"
			}
		} else {
			s += "\n"
			for (let line of sel.split("\n")) {
				s += ">" + line + "\n"
			}
		}

		// Don't commit a quote, if it is the only input in a post
		this.view.replaceText(old + s,
			postSM.state !== postState.draft || old.length !== 0)
	}

	// Commit a post made with live updates disabled
	public async commitNonLive() {
		let files: FileList
		if (this.view.upload) {
			files = this.view.upload.input.files
		}
		const text = this.view.input.value;
		if (!text.length && !files.length) {
			return postSM.feed(postEvent.done);
		}

		const req = newAllocRequest()
		if (this.view.upload && files.length) {
			req["image"] = await this.view.upload.uploadFile(files[0])
		}
		if (text) {
			req["body"] = text;
		}
		send(message.insertPost, req);
		handlers[message.postID] = this.receiveID();
	}

	// Returns a function, that handles a message from the server, containing
	// the ID of the allocated post.
	private receiveID(): (id: number) => void {
		return (id: number) => {
			this.id = id
			this.op = page.thread
			this.seenOnce = true
			postSM.feed(postEvent.alloc)
			storeSeenPost(this.id, this.op)
			storeMine(this.id, this.op)
			posts.add(this)
			delete handlers[message.postID]
		}
	}

	// Request allocation of a draft post to the server
	private requestAlloc(body: string = this.view.input.value,
		image: FileData | null,
	) {
		const req = newAllocRequest();
		req["open"] = true;
		if (body) {
			req["body"] = this.inputBody = body;
		}
		if (image) {
			req["image"] = image;
		}

		send(message.insertPost, req);
		postSM.feed(postEvent.sentAllocRequest);
		handlers[message.postID] = this.receiveID();
	}

	// Handle draft post allocation
	public onAllocation(data: PostData) {
		extend(this, data);
		if (postSM.state === postState.alloc) {
			this.view.renderAlloc();
			if (this.image) {
				this.insertImage(this.image);
			}
		} else {
			this.propagateLinks();
		}
	}

	// Upload the file and request its allocation
	public async uploadFile(files?: FileList | File) {
		if (boardConfig.textOnly) {
			return
		}
		if (files && this.view.upload) {
			// Can't set value with an individual File
			if (files instanceof FileList) {
				(this.view.upload.input.files as any) = files;
			}
		}

		// Already have image or not in live mode
		if (this.image || !identity.live) {
			return
		}

		const data = await this.view.upload
			.uploadFile(files instanceof FileList ? files[0] : files);
		// Upload failed, canceled, image added while thumbnailing or post
		// closed
		if (!data || this.image || !this.editing) {
			return
		}

		if (postSM.state === postState.draft) {
			this.requestAlloc(null, data)
		} else {
			send(message.insertImage, data)
		}
	}

	// Insert the uploaded image into the model
	public insertImage(img: ImageData) {
		this.image = img
		this.view.insertImage()
	}

	// Spoiler an already allocated image
	public commitSpoiler() {
		this.send(message.spoiler, null)
	}
}

// Find the first differing character in 2 character arrays
function diffIndex(a: string[], b: string[]): number {
	for (let i = 0; i < a.length; i++) {
		if (a[i] !== b[i]) {
			return i
		}
	}
	return a.length
}
