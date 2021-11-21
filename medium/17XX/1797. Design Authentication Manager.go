package _7XX

type AuthenticationManager struct {
	ttl       int
	startTime map[string]int
	endTime   map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{ttl: timeToLive, startTime: map[string]int{}, endTime: map[string]int{}}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.startTime[tokenId] = currentTime
	this.endTime[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if this.endTime[tokenId] <= currentTime{
		delete(this.startTime,tokenId)
		delete(this.endTime,tokenId)

		return
	}

	this.startTime[tokenId] = currentTime
	this.endTime[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	res := 0

	for k, v := range this.startTime {
		if currentTime >= v && currentTime < this.endTime[k] {
			res++
		}
	}

	return res
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
