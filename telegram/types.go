package telegram

type ChatID any

type User struct {
	ID                      int64   `json:"id"`
	IsBot                   bool    `json:"is_bot"`
	FirstName               string  `json:"first_name"`
	LastName                *string `json:"last_name,omitempty"`
	Username                *string `json:"username,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	IsPremium               *bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   *bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"` // Corrected
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness    *bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           *bool   `json:"has_main_web_app,omitempty"`
}

type Chat struct {
	ID        int64   `json:"id"`
	Type      string  `json:"type"`
	Title     *string `json:"title,omitempty"`
	Username  *string `json:"username,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	IsForum   *bool   `json:"is_forum,omitempty"`
}

type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`
	Type                               string                `json:"type"`
	Title                              *string               `json:"title,omitempty"`
	Username                           *string               `json:"username,omitempty"`
	FirstName                          *string               `json:"first_name,omitempty"`
	LastName                           *string               `json:"last_name,omitempty"`
	IsForum                            *bool                 `json:"is_forum,omitempty"`
	AccentColorID                      int                   `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiID            *string               `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorID               *int                  `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiID     *string               `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiID           *string               `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          *int                  `json:"emoji_status_expiration_date,omitempty"`
	BIO                                *string               `json:"bio,omitempty"`
	HasPrivateForwards                 *bool                 `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages *bool                 `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 *bool                 `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      *bool                 `json:"join_by_request,omitempty"`
	Description                        *string               `json:"description,omitempty"`
	InviteLink                         *string               `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	AcceptedGiftTypes                  AcceptedGiftTypes     `json:"accepted_gift_types"`
	CanSendPaidMedia                   *bool                 `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      *int                  `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               *int                  `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              *int                  `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       *bool                 `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   *bool                 `json:"has_hidden_members,omitempty"`
	HasProtectedContent                *bool                 `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  *bool                 `json:"has_visible_history,omitempty"`
	StickerSetName                     *string               `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   *bool                 `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          *string               `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatId                       *int64                `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
}

type Message struct {
	MessageID                     int                            `json:"message_id"`
	MessageThreadID               *int                           `json:"message_thread_id,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	SenderBoostCount              *int                           `json:"sender_boost_count,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	Date                          int                            `json:"date"`
	BusinessConnectionID          *string                        `json:"business_connection_id,omitempty"`
	Chat                          Chat                           `json:"chat"`
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`
	IsTopicMessage                *bool                          `json:"is_topic_message,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      *int                           `json:"edit_date,omitempty"`
	HasProtectedContent           *bool                          `json:"has_protected_content,omitempty"`
	IsFromOffline                 *bool                          `json:"is_from_offline,omitempty"`
	MediaGroupID                  *string                        `json:"media_group_id,omitempty"`
	AuthorSignature               *string                        `json:"author_signature,omitempty"`
	PaidStarCount                 *int                           `json:"paid_star_count,omitempty"`
	Text                          *string                        `json:"text,omitempty"`
	Entities                      []MessageEntity                `json:"entities,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	EffectID                      *string                        `json:"effect_id,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`
	Photo                         []PhotoSize                    `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       *string                        `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia         *bool                          `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler               *bool                          `json:"has_media_spoiler,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  *string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               *bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              *bool                          `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         *bool                          `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            *bool                          `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               *int64                         `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             *int64                         `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	Gift                          *GiftInfo                      `json:"gift,omitempty"`
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift,omitempty"`
	ConnectedWebsite              *string                        `json:"connected_website,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`
	PaidMessagePriceChanged       *PaidMessagePriceChanged       `json:"paid_message_price_changed,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

type MessageID struct {
	MessageID int64 `json:"message_id"`
}

type InaccessibleMessage struct {
	Chat      Chat  `json:"chat"`
	MessageID int64 `json:"message_id"`
	Date      int   `json:"date"`
}

type MaybeInaccessibleMessage any

type MessageEntity struct {
	Type          string  `json:"type"`
	Offset        int     `json:"offset"`
	Length        int     `json:"length"`
	URL           *string `json:"url,omitempty"`
	User          *User   `json:"user,omitempty"`
	Language      *string `json:"language,omitempty"`
	CustomEmojiID *string `json:"custom_emoji_id,omitempty"`
}

type TextQuote struct {
	Text     string          `json:"text"`
	Entities []MessageEntity `json:"entities,omitempty"`
	Position int             `json:"position"`
	IsManual *bool           `json:"is_manual,omitempty"`
}

type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`
	Chat               *Chat               `json:"chat,omitempty"`
	MessageID          *int64              `json:"message_id,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation          *Animation          `json:"animation,omitempty"`
	Audio              *Audio              `json:"audio,omitempty"`
	Document           *Document           `json:"document,omitempty"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`
	Photo              []PhotoSize         `json:"photo,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
	Story              *Story              `json:"story,omitempty"`
	Video              *Video              `json:"video,omitempty"`
	VideoNote          *VideoNote          `json:"video_note,omitempty"`
	Voice              *Voice              `json:"voice,omitempty"`
	HasMediaSpoiler    *bool               `json:"has_media_spoiler,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
	Dice               *Dice               `json:"dice,omitempty"`
	Game               *Game               `json:"game,omitempty"`
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`
	Invoice            *Invoice            `json:"invoice,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
}

type ReplyParameters struct {
	MessageID                int64           `json:"message_id"`
	ChatId                   *ChatID         `json:"chat_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	Quote                    *string         `json:"quote,omitempty"`
	QuoteParseMode           *string         `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            *int            `json:"quote_position,omitempty"`
}

type MessageOrigin any

type MessageOriginUser struct {
	Type       string `json:"type"`
	Date       int    `json:"date"`
	SenderUser User   `json:"sender_user"`
}

type MessageOriginHiddenUser struct {
	Type           string `json:"type"`
	Date           int    `json:"date"`
	SenderUserName string `json:"sender_user_name"`
}

type MessageOriginChat struct {
	Type            string  `json:"type"`
	Date            int     `json:"date"`
	SenderChat      Chat    `json:"sender_chat"`
	AuthorSignature *string `json:"author_signature,omitempty"`
}

type MessageOriginChannel struct {
	Type            string  `json:"type"`
	Date            int     `json:"date"`
	Chat            Chat    `json:"chat"`
	MessageId       int64   `json:"message_id"`
	AuthorSignature *string `json:"author_signature,omitempty"`
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     *int   `json:"file_size,omitempty"`
}

type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    *string    `json:"performer,omitempty"`
	Title        *string    `json:"title,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int64     `json:"file_size,omitempty"`
}

type Story struct {
	Chat Chat  `json:"chat"`
	ID   int64 `json:"id"`
}

type Video struct {
	FileID         string      `json:"file_id"`
	FileUniqueID   string      `json:"file_unique_id"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Duration       int         `json:"duration"`
	Thumbnail      *PhotoSize  `json:"thumbnail,omitempty"`
	Cover          []PhotoSize `json:"cover,omitempty"`
	StartTimestamp *int        `json:"start_timestamp,omitempty"`
	FileName       *string     `json:"file_name,omitempty"`
	MimeType       *string     `json:"mime_type,omitempty"`
	FileSize       *int64      `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
}

type Voice struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	Duration     int     `json:"duration"`
	MimeType     *string `json:"mime_type,omitempty"`
	FileSize     *int64  `json:"file_size,omitempty"`
}

type PaidMediaInfo struct {
	StarCount int         `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

type PaidMedia any

type PaidMediaPreview struct {
	Type     string `json:"type"`
	Width    *int   `json:"width,omitempty"`
	Height   *int   `json:"height,omitempty"`
	Duration *int   `json:"duration,omitempty"`
}

type PaidMediaPhoto struct {
	Type  string      `json:"type"`
	Photo []PhotoSize `json:"photo"`
}

type PaidMediaVideo struct {
	Type  string `json:"type"`
	Video Video  `json:"video"`
}

type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int64  `json:"user_id,omitempty"`
	Vcard       *string `json:"vcard,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type PollOption struct {
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int             `json:"voter_count"`
}

type InputPollOption struct {
	Text          string          `json:"text"`
	TextParseMode *string         `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

type PollAnswer struct {
	PollID    string `json:"poll_id"`
	VoterChat *Chat  `json:"voter_chat,omitempty"`
	User      *User  `json:"user,omitempty"`
	OptionIDs []int  `json:"option_ids"`
}

type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	QuestionEntities      []MessageEntity `json:"question_entities,omitempty"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       *int            `json:"correct_option_id,omitempty"`
	Explanation           *string         `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            *int            `json:"open_period,omitempty"`
	CloseDate             *int            `json:"close_date,omitempty"`
}

type Location struct {
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int     `json:"live_period,omitempty"`
	Heading              *int     `json:"heading,omitempty"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"`
}

type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareID    *string  `json:"foursquare_id,omitempty"`
	FoursquareType  *string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string  `json:"google_place_id,omitempty"`
	GooglePlaceType *string  `json:"google_place_type,omitempty"`
}

type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

type BackgroundFill any

type BackgroundFillSolid struct {
	Type  string `json:"type"`
	Color int    `json:"color"`
}

type BackgroundFillGradient struct {
	Type          string `json:"type"`
	TopColor      int    `json:"top_color"`
	BottomColor   int    `json:"bottom_color"`
	RotationAngle int    `json:"rotation_angle"`
}

type BackgroundFillFreeformGradient struct {
	Type   string `json:"type"`
	Colors []int  `json:"colors"`
}

type BackgroundType any

type BackgroundTypeFill struct {
	Type             string         `json:"type"`
	Fill             BackgroundFill `json:"fill"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
}

type BackgroundTypePattern struct {
	Type       string         `json:"type"`
	Document   Document       `json:"document"`
	Fill       BackgroundFill `json:"fill"`
	Intensity  int            `json:"intensity"`
	IsInverted *bool          `json:"is_inverted,omitempty"`
	IsMoving   *bool          `json:"is_moving,omitempty"`
}

type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}

type ChatBackground struct {
	Type BackgroundType `json:"type"`
}

type ForumTopicCreated struct {
	Name              string  `json:"name"`
	IconColor         int     `json:"icon_color"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicClosed struct{}

type ForumTopicEdited struct {
	Name              *string `json:"name,omitempty"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicReopened struct{}

type GeneralForumTopicHidden struct{}

type GeneralForumTopicUnhidden struct{}

type SharedUser struct {
	UserID    int64       `json:"user_id"`
	FirstName *string     `json:"first_name,omitempty"`
	LastName  *string     `json:"last_name,omitempty"`
	Username  *string     `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

type UsersShared struct {
	RequestID int          `json:"request_id"`
	User      []SharedUser `json:"user"`
}

type ChatShared struct {
	RequestID int         `json:"request_id"`
	ChatID    int64       `json:"chat_id"`
	Title     *string     `json:"title,omitempty"`
	Username  *string     `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

type WriteAccessAllowed struct {
	FromRequest        *bool   `json:"from_request,omitempty"`
	WebAppName         *string `json:"web_app_name,omitempty"`
	FromAttachmentMenu *bool   `json:"from_attachment_menu,omitempty"`
}

type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

type VideoChatStarted struct{}

type VideoChatEnded struct {
	Duration int `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

type GiveawayCreated struct {
	PrizeStarCount *int `json:"prize_star_count,omitempty"`
}

type Giveaway struct {
	Chats                         []Chat   `json:"chats"`
	WinnersSelectionDate          int      `json:"winners_selection_date"`
	WinnerCount                   int      `json:"winner_count"`
	OnlyNewMembers                *bool    `json:"only_new_members,omitempty"`
	HasPublicWinners              *bool    `json:"has_public_winners,omitempty"`
	PrizeDescription              *string  `json:"prize_description,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	PrizeStarCount                *int     `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount *int     `json:"premium_subscription_month_count,omitempty"`
}

type GiveawayWinners struct {
	Chat                          Chat    `json:"chat"`
	GiveawayMessageID             int64   `json:"giveaway_message_id"`
	WinnersSelectionDate          int     `json:"winners_selection_date"`
	WinnerCount                   int     `json:"winner_count"`
	Winners                       []User  `json:"winners"`
	AdditionalChatCount           *int    `json:"additional_chat_count,omitempty"`
	PrizeStarCount                *int    `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount *int    `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           *int    `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                *bool   `json:"only_new_members,omitempty"`
	WasRefunded                   *bool   `json:"was_refunded,omitempty"`
	PrizeDescription              *string `json:"prize_description,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount *int     `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	IsStarGiveaway      *bool    `json:"is_star_giveaway,omitempty"`
}

type LinkPreviewOptions struct {
	IsDisabled       *bool   `json:"is_disabled,omitempty"`
	URL              *string `json:"url,omitempty"`
	PreferSmallMedia *bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia *bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    *bool   `json:"show_above_text,omitempty"`
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type File struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	FileSize     *int    `json:"file_size,omitempty"`
	FilePath     *string `json:"file_path,omitempty"`
}

type WebAppInfo struct {
	URL string `json:"url"`
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          *bool              `json:"is_persistent ,omitempty"`
	ResizeKeyboard        *bool              `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       *bool              `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder *string            `json:"input_field_placeholder,omitempty"`
	Selective             *bool              `json:"selective,omitempty"`
}

type KeyboardButton struct {
	Text            string                      `json:"text"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact  *bool                       `json:"request_contact,omitempty"`
	RequestLocation *bool                       `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
}

type KeyboardButtonRequestUsers struct {
	RequestID       int   `json:"request_id"`
	UserIsBot       *bool `json:"user_is_bot,omitempty"`
	UserIsPremium   *bool `json:"user_is_premium,omitempty"`
	MaxQuantity     *int  `json:"max_quantity,omitempty"`
	RequestName     *bool `json:"request_name,omitempty"`
	RequestUsername *bool `json:"request_username,omitempty"`
	RequestPhoto    *bool `json:"request_photo,omitempty"`
}

type KeyboardButtonRequestChat struct {
	RequesID                int                      `json:"request_id"`
	ChatIsChannel           bool                     `json:"chat_is_channel"`
	ChatIsForum             *bool                    `json:"chat_is_forum,omitempty"`
	ChatHasUsername         *bool                    `json:"chat_has_username,omitempty"`
	ChatIsCreated           *bool                    `json:"chat_is_created,omitempty"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotIsMember             *bool                    `json:"bot_is_member,omitempty"`
	RequestTitle            *bool                    `json:"request_title,omitempty"`
	RequestUsername         *bool                    `json:"request_username,omitempty"`
	RequestPhoto            *bool                    `json:"request_photo,omitempty"`
}

type KeyboardButtonPollType struct {
	Type *string `json:"type,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	URL                          *string                      `json:"url,omitempty"`
	CallbackData                 *string                      `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginURL                     *LoginURL                    `json:"login_url,omitempty"`
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          *bool                        `json:"pay,omitempty"`
}

// TUT<===============================================================================
type LoginURL struct {
	URL                string  `json:"url"`
	ForwardText        *string `json:"forward_text,omitempty"`
	BotUsername        *string `json:"bot_username,omitempty"`
	RequestWriteAccess *bool   `json:"request_write_access,omitempty"`
}

type SwitchInlineQueryChosenChat struct {
	Query             *string `json:"query,omitempty"`
	AllowUserChats    *bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     *bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   *bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats *bool   `json:"allow_channel_chats,omitempty"`
}

type CopyTextButton struct {
	Text string `json:"text"`
}

type CallbackQuery struct {
	ID              string                    `json:"id"`
	From            *User                     `json:"from"`
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`
	InlineMessageID *string                   `json:"inline_message_id,omitempty"`
	ChatInstance    string                    `json:"chat_instance"`
	Data            *string                   `json:"data,omitempty"`
	GameShortName   *string                   `json:"game_short_name,omitempty"`
}

type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
	Selective             *bool   `json:"selective,omitempty"`
}

type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`
	Creator                 *User   `json:"creator"`
	CreatesJoinRequest      bool    `json:"creates_join_request"`
	IsPrimary               bool    `json:"is_primary"`
	IsRevoked               bool    `json:"is_revoked"`
	Name                    *string `json:"name,omitempty"`
	ExpireDate              *int    `json:"expire_date,omitempty"`
	MemberLimit             *int    `json:"member_limit,omitempty"`
	PendingJoinRequestCount *int    `json:"pending_join_request_count,omitempty"`
	SubscriptionPeriod      *int    `json:"subscription_period,omitempty"`
	SubscriptionPrice       *int    `json:"subscription_price,omitempty"`
}

type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`
	CanManageChat       bool `json:"can_manage_chat"`
	CanDeleteMessages   bool `json:"can_delete_messages"`
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	CanRestrictMembers  bool `json:"can_restrict_members"`
	CanPromoteMembers   bool `json:"can_promote_members"`
	CanChangeInfo       bool `json:"can_change_info"`
	CanInviteUsers      bool `json:"can_invite_users"`
	CanPostStories      bool `json:"can_post_stories"`
	CanEditStories      bool `json:"can_edit_stories"`
	CanDeleteStories    bool `json:"can_delete_stories"`
	CanPostMessages     bool `json:"can_post_messages"`
	CanEditMessages     bool `json:"can_edit_messages"`
	CanPinMessages      bool `json:"can_pin_messages"`
	CanManageTopics     bool `json:"can_manage_topics"`
}

type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
	Date                    int             `json:"date"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          *bool           `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink *bool           `json:"via_chat_folder_invite_link,omitempty"`
}

type ChatMember any

type ChatMemberAdministrator struct {
	Status              string  `json:"status"`
	User                User    `json:"user"`
	CanBeEdited         bool    `json:"can_be_edited"`
	IsAnonymous         bool    `json:"is_anonymous"`
	CanManageChat       bool    `json:"can_manage_chat"`
	CanDeleteMessages   bool    `json:"can_delete_messages"`
	CanManageVideoChats bool    `json:"can_manage_video_chats"`
	CanRestrictMembers  bool    `json:"can_restrict_members"`
	CanPromoteMembers   bool    `json:"can_promote_members"`
	CanChangeInfo       bool    `json:"can_change_info"`
	CanInviteUsers      bool    `json:"can_invite_users"`
	CanPostStories      bool    `json:"can_post_stories"`
	CanEditStories      bool    `json:"can_edit_stories"`
	CanDeleteStories    bool    `json:"can_delete_stories"`
	CanPostMessages     *bool   `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool   `json:"can_pin_messages,omitempty"`
	CanManageTopics     *bool   `json:"can_manage_topics,omitempty"`
	CustomTitle         *string `json:"custom_title,omitempty"`
}

type ChatMemberMember struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate *int   `json:"until_date,omitempty"`
}

type ChatMemberRestricted struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendAudios         bool   `json:"can_send_audios"`
	CanSendDocuments      bool   `json:"can_send_documents"`
	CanSendPhotos         bool   `json:"can_send_photos"`
	CanSendVideos         bool   `json:"can_send_videos"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`
	CanSenVoiceNotes      bool   `json:"can_sen_voice_notes"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanManageTopics       bool   `json:"can_manage_topics"`
	UntilDate             int    `json:"until_date"`
}

type ChatMemberLeft struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate int    `json:"until_date"`
}

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       *User           `json:"from"`
	UserChatID int64           `json:"user_chat_id"`
	Date       int             `json:"date"`
	BIO        *string         `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatPermissions struct {
	CanSendMessages       *bool `json:"can_send_messages,omitempty"`
	CanSendAudios         *bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      *bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         *bool `json:"can_send_photos,omitempty"`
	CanSendVideos         *bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     *bool `json:"can_send_video_notes,omitempty"`
	CansSendVoiceNotes    *bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       *bool `json:"can_manage_topics,omitempty"`
}
type Birthdate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type BusinessIntro struct {
	Title   *string `json:"title,omitempty"`
	Message *string `json:"message,omitempty"`
	Sticker *string `json:"sticker,omitempty"`
}

type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

type ReactionType any

type ReactionTypeEmoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"`
	CustomEmojiID string `json:"custom_emoji_id"`
}

type ReactionTypePaid struct {
	Type string `json:"type"`
}

type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	MessageID   int            `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int            `json:"date"`
	OldReaction []ReactionType `json:"old_reaction,omitempty"`
	NewReaction []ReactionType `json:"new_reaction,omitempty"`
}

type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int             `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

type ForumTopic struct {
	MessageThreadId   int     `json:"message_thread_id"`
	Name              string  `json:"name"`
	IconColor         int     `json:"icon_color"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type BotCommandScope any

type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"`
}

type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"`
}

type BotCommandScopeAllChatAdministrators struct {
	Type   string `json:"type"`
	ChatID ChatID `json:"chat_id"`
}

type BotCommandScopeChat struct {
	Type   string `json:"type"`
	ChatID ChatID `json:"chat_id"`
}

type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`
	ChatID ChatID `json:"chat_id"`
}

type BotCommandScopeChatMember struct {
	Type   string `json:"type"`
	ChatID ChatID `json:"chat_id"`
	UserID int    `json:"user_id"`
}

type BotName struct {
	Name string `json:"name"`
}

type BotDescription struct {
	Description string `json:"description"`
}

type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

type MenuButton any

type MenuButtonCommands struct {
	Type string `json:"type"`
}

type MenuButtonWebApp struct {
	Type   string     `json:"type"`
	Text   string     `json:"text"`
	WebApp WebAppInfo `json:"web_app"`
}

type MenuButtonDefault struct {
	Type string `json:"type"`
}

type ChatBoostSource any

type ChatBoostSourcePremium struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

type ChatBoostSourceGiftCode struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`
	GiveawayMessageID int    `json:"giveaway_message_id"`
	User              *User  `json:"user,omitempty"`
	PrizeStarCount    *int   `json:"prize_star_count,omitempty"`
	IsUnclaimed       *bool  `json:"is_unclaimed,omitempty"`
}

type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int             `json:"add_date"`
	ExpirationDate int             `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int             `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}

type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

type BusinessBotRights struct {
	CanReply                   *bool `json:"can_reply,omitempty"`
	CanReadMessages            *bool `json:"can_read_messages,omitempty"`
	CanDeleteOutgoingMessages  *bool `json:"can_delete_outgoing_messages,omitempty"`
	CanDeleteAllMessages       *bool `json:"can_delete_all_messages,omitempty"`
	CanEditMessages            *bool `json:"can_edit_messages,omitempty"`
	CanEditName                *bool `json:"can_edit_name,omitempty"`
	CanEditBIO                 *bool `json:"can_edit_bio,omitempty"`
	CanEditProfilePhoto        *bool `json:"can_edit_profile_photo,omitempty"`
	CanEditUsername            *bool `json:"can_edit_username,omitempty"`
	CanChangeGiftSettings      *bool `json:"can_change_gift_settings,omitempty"`
	CanViewGiftsAndStars       *bool `json:"can_view_gifts_and_stars,omitempty"`
	CanConvertGiftsToStars     *bool `json:"can_convert_gifts_to_stars,omitempty"`
	CanTransferAndUpgradeGifts *bool `json:"can_transfer_and_upgrade_gifts,omitempty"`
	CanTransferStars           *bool `json:"can_transfer_stars,omitempty"`
	CanManageStories           *bool `json:"can_manage_stories,omitempty"`
}

type BusinessConnection struct {
	ID         string             `json:"id"`
	User       User               `json:"user"`
	UserChatID int64              `json:"user_chat_id"`
	Date       int                `json:"date"`
	Rights     *BusinessBotRights `json:"rights,omitempty"`
	IsEnabled  bool               `json:"is_enabled"`
}

type BusinessMessagesDeleted struct {
	BusinessConnectionID string  `json:"business_connection_id"`
	Chat                 Chat    `json:"chat"`
	MessageIDs           []int64 `json:"message_ids"`
}

type ResponseParameters struct {
	MigrateToChatID *int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int   `json:"retry_after,omitempty"`
}

type InputMedia any

type InputMediaPhoto struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

type InputMediaVideo struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Thumbnail             *string         `json:"thumbnail,omitempty"`
	Cover                 *string         `json:"cover,omitempty"`
	StartTimestamp        *int            `json:"start_timestamp,omitempty"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	Width                 *int            `json:"width,omitempty"`
	Height                *int            `json:"height,omitempty"`
	Duration              *int            `json:"duration,omitempty"`
	SupportsStreaming     *bool           `json:"supports_streaming,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

type InputMediaAnimation struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Thumbnail             *string         `json:"thumbnail,omitempty"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	Width                 *int            `json:"width,omitempty"`
	Height                *int            `json:"height,omitempty"`
	Duration              *int            `json:"duration,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

type InputMediaAudio struct {
	Type            string          `json:"type"`
	Media           string          `json:"media"`
	Thumbnail       *string         `json:"thumbnail,omitempty"`
	Caption         *string         `json:"caption,omitempty"`
	ParseMode       *string         `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        *int            `json:"duration,omitempty"`
	Performer       *string         `json:"performer,omitempty"`
	Title           *string         `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Type                        string          `json:"type"`
	Media                       string          `json:"media"`
	Thumbnail                   *string         `json:"thumbnail,omitempty"`
	Caption                     *string         `json:"caption,omitempty"`
	ParseMode                   *string         `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection *bool           `json:"disable_content_type_detection,omitempty"`
}

type InputFile struct{}

type InputPaidMedia any

type InputPaidMediaPhoto struct {
	Type  string `json:"type"`
	Media string `json:"media"`
}

type InputPaidMediaVideo struct {
	Type              string  `json:"type"`
	Media             string  `json:"media"`
	Thumbnail         *string `json:"thumbnail,omitempty"`
	Cover             *string `json:"cover,omitempty"`
	StartTimestamp    *int    `json:"start_timestamp,omitempty"`
	Width             *int    `json:"width,omitempty"`
	Height            *int    `json:"height,omitempty"`
	Duration          *int    `json:"duration,omitempty"`
	SupportsStreaming *bool   `json:"supports_streaming,omitempty"`
}

type InputProfilePhoto any

type InputProfilePhotoStatic struct {
	Type  string `json:"type"`
	Photo string `json:"photo"`
}

type InputProfilePhotoAnimated struct {
	Type               string   `json:"type"`
	Animation          string   `json:"animation"`
	MainFrameTimestamp *float64 `json:"main_frame_timestamp,omitempty"`
}

type InputStoryContent any

type InputStoryContentPhoto struct {
	Type  string `json:"type"`
	Photo string `json:"photo"`
}

type InputStoryContentVideo struct {
	Type                string   `json:"type"`
	Video               string   `json:"video"`
	Duration            *float64 `json:"duration,omitempty"`
	CoverFrameTimestamp *float64 `json:"cover_frame_timestamp,omitempty"`
	IsAnimation         *bool    `json:"is_animation,omitempty"`
}

//-------------------------------------------------------------

//-------------------------------------------------------------
//GAME TYPES

type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         *string         `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

type CallbackGame struct{}

type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}
