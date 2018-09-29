package models

import (
	"database/sql"
)

type Topic struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	TextSource string `json:"text_source"`
	TextHash   string `json:"text_hash"`
}

func GetAllTopics(db *sql.DB) (topics []Topic, err error) {
	rows, err := db.Query("SELECT topic_id, topic_title, topic_text, topic_text_source, topic_text_hash FROM topics")
	for rows.Next() {
		var topic Topic
		if ok := rows.Scan(&topic.ID, &topic.Title, &topic.Text, &topic.TextSource, &topic.TextHash); ok == nil {
			topics = append(topics, topic)
		}
	}
	return topics, err
}

func GetTopicByID(id int, db *sql.DB) (Topic, error) {
	var topic Topic
	err := db.QueryRow("SELECT topic_id, topic_title, topic_text, topic_text_source, topic_text_hash "+
		"FROM topics WHERE topic_id = $1", id).Scan(&topic.ID, &topic.Title, &topic.Text, &topic.TextSource, &topic.TextHash)
	return topic, err
}

func (topic *Topic) Save(db *sql.DB) (topicID int, err error) {
	err = db.QueryRow("INSERT INTO topics(topic_id, topic_title, topic_text, topic_text_source, topic_text_hash) VALUES($1, $2, $3, $4, $5) "+
		"ON CONFLICT(topic_id) DO UPDATE SET topic_title = $2, topic_text = $3, topic_text_source = $4, topic_text_hash = $5 "+
		"RETURNING topic_id", topic.ID, topic.Title, topic.Text, topic.Text, "123").Scan(&topicID)
	return topicID, err
}
