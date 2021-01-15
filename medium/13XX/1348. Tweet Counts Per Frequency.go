package t13XX


type TweetCounts struct {
	tweetMap map[string][]int
}

func TweetCountsConstructor() TweetCounts {
	t := TweetCounts{
		tweetMap:map[string][]int{},
	}


	return t
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.tweetMap[tweetName] = append(this.tweetMap[tweetName], time)
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {

	multiplier := 60
	switch freq {
	case "hour":
		multiplier = 3600
	case "day":
		multiplier = 86400
	}

	result := make([]int,(endTime - startTime) / multiplier+1)

	for _, v := range this.tweetMap[tweetName] {
		if v >= startTime && v <= endTime {
			result[(v - startTime) / multiplier] += 1
		}
	}

	return result
}


/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
