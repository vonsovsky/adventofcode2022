package main

type segment struct {
	start, end int
}

var segments []segment

func GetSegments() []segment {
	return segments
}

func ResetSegments() {
	segments = []segment{}
}

func insertAtIndex(i int, element segment) {
	segments = append(segments[:i], append([]segment{element}, segments[i:]...)...)
}

func removeAtIndex(i int) {
	segments = append(segments[:i], segments[i+1:]...)
}

func consolidate(prev int, next int) {
	segments[prev].end = max(segments[prev].end, segments[next].end)
	removeAtIndex(next)
}

func AddSegment(start int, end int) {
	insertAt := len(segments)
	for i, s := range segments {
		if start < s.start {
			insertAt = i
			break
		}
	}
	insertAtIndex(insertAt, segment{
		start: start,
		end:   end,
	})

	for insertAt > 0 && segments[insertAt-1].end+1 >= segments[insertAt].start {
		consolidate(insertAt-1, insertAt)
		insertAt -= 1
	}

	for insertAt < len(segments)-1 && segments[insertAt+1].start <= segments[insertAt].end+1 {
		consolidate(insertAt, insertAt+1)
	}
}
