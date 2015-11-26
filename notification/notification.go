package notification

import "time"

const (
  High = 3
  Medium = 2
  Low = 1
)

type Notification struct{
  ID uint64
  Title string
  Tag1 string
  Tag2 string
  Tag3 string
  Info string
  Relevance uint8
  Transaction string
  TransactionResult bool
  Company uint
  DateTime time.Time
}
