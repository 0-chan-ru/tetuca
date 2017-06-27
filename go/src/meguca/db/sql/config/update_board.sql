update boards
	set
		readOnly = $2,
		textOnly = $3,
		forcedAnon = $4,
		disableRobots = $5,
		defaultCSS = $6,
		title = $7,
		notice = $8,
		rules = $9,
		eightball = $10
	where id = $1
	returning pg_notify('board_updated', $1)
