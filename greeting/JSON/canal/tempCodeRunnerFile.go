if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	