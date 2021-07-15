package redis

// redis key

const (
	KeyPrefix        = "web:"
	KeyPostTimeZSet  = "post:time"     // ZSet 帖子及发帖时间
	KeyPostScoreZSet = "post:score"    //  ZSet 帖子及投票的分数
	KeyPostVoteZSet  = "post:voted:%d" //  ZSet 记录用户的投票记录，%d为用户的ID
)
