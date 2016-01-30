package main

import (
	"fmt"
	"time"

	"github.com/mediocregopher/radix.v2/redis"
)

const (
	ONE_WEEK_IN_SECONDS = 7 * 86400
	VOTE_SCORE = 432 // 86400 / 200
	ARTICLES_PER_PAGE = 25
)

func article_vote(conn, user, article) {
	// Calculate the cutoff time for voting
	cutoff = time.time() - ONE_WEEK_IN_SECONDS
}

func post_article() {

}

func get_articles () {

}

func add_remove_groups () {

}

func get_group_articles () {

}

func main() {

}