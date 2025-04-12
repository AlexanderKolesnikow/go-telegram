package telegram

//  AVAILABLE TYPES https://core.telegram.org/bots/api#available-types

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

type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
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
	From            User                      `json:"from"`
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
	Creator                 User    `json:"creator"`
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
	IsAnonymous         bool  `json:"is_anonymous"`
	CanManageChat       bool  `json:"can_manage_chat"`
	CanDeleteMessages   bool  `json:"can_delete_messages"`
	CanManageVideoChats bool  `json:"can_manage_video_chats"`
	CanRestrictMembers  bool  `json:"can_restrict_members"`
	CanPromoteMembers   bool  `json:"can_promote_members"`
	CanChangeInfo       bool  `json:"can_change_info"`
	CanInviteUsers      bool  `json:"can_invite_users"`
	CanPostStories      bool  `json:"can_post_stories"`
	CanEditStories      bool  `json:"can_edit_stories"`
	CanDeleteStories    bool  `json:"can_delete_stories"`
	CanPostMessages     *bool `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics     *bool `json:"can_manage_topics,omitempty"`
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
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`
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
	From       User            `json:"from"`
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
	CanSendVoiceNotes     *bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       *bool `json:"can_manage_topics,omitempty"`
}
type Birthdate struct {
	Day   int  `json:"day"`
	Month int  `json:"month"`
	Year  *int `json:"year,omitempty"`
}

type BusinessIntro struct {
	Title   *string  `json:"title,omitempty"`
	Message *string  `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
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

type StoryAreaPosition struct {
	XPercentage            float64 `json:"x_percentage"`
	YPercentage            float64 `json:"y_percentage"`
	WidthPercentage        float64 `json:"width_percentage"`
	HeightPercentage       float64 `json:"height_percentage"`
	RotationAngle          float64 `json:"rotation_angle"`
	CornerRadiusPercentage float64 `json:"corner_radius_percentage"`
}

type LocationAddress struct {
	CountryCode string  `json:"country_code"`
	State       *string `json:"state,omitempty"`
	City        *string `json:"city,omitempty"`
	Street      *string `json:"street,omitempty"`
}

type StoryAreaType any

type StoryAreaTypeLocation struct {
	Type      string           `json:"type"`
	Latitude  float64          `json:"latitude"`
	Longitude float64          `json:"longitude"`
	Address   *LocationAddress `json:"address,omitempty"`
}

type StoryAreaTypeSuggestedReaction struct {
	Type         string       `json:"type"`
	ReactionType ReactionType `json:"reaction_type"`
	IsDark       *bool        `json:"is_dark,omitempty"`
	IsFlipped    *bool        `json:"is_flipped,omitempty"`
}

type StoryAreaTypeLink struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type StoryAreaTypeWeather struct {
	Type            string  `json:"type"`
	Temperature     float64 `json:"temperature"`
	Emoji           string  `json:"emoji"`
	BackgroundColor int     `json:"background_color"`
}

type StoryAreaTypeUniqueGift struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type StoryArea struct {
	Position StoryAreaPosition `json:"position"`
	Type     StoryAreaType     `json:"type"`
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
	MessageID   int64          `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int            `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageID int64           `json:"message_id"`
	Date      int             `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

type ForumTopic struct {
	MessageThreadId   int     `json:"message_thread_id"`
	Name              string  `json:"name"`
	IconColor         int     `json:"icon_color"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}

type Gift struct {
	ID               string  `json:"id"`
	Sticker          Sticker `json:"sticker"`
	StarCount        int     `json:"star_count"`
	UpgradeStarCount *int    `json:"upgrade_star_count,omitempty"`
	TotalCount       *int    `json:"total_count,omitempty"`
	RemainingCount   *int    `json:"remaining_count,omitempty"`
}

type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

type UniqueGiftModel struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
}

type UniqueGiftSymbol struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
}

type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	SymbolColor int `json:"symbol_color"`
	TextColor   int `json:"text_color"`
}

type UniqueGiftBackdrop struct {
	Name           string                   `json:"name"`
	Colors         UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                      `json:"rarity_per_mille"`
}

type UniqueGift struct {
	BaseName string             `json:"base_name"`
	Name     string             `json:"name"`
	Number   int                `json:"number"`
	Model    UniqueGiftModel    `json:"model"`
	Symbol   UniqueGiftSymbol   `json:"symbol"`
	Backdrop UniqueGiftBackdrop `json:"backdrop"`
}

type GiftInfo struct {
	Gift                    Gift            `json:"gift"`
	OwnedGiftID             *string         `json:"owned_gift_id,omitempty"`
	ConvertStarCount        *int            `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount *int            `json:"prepaid_upgrade_star_count,omitempty"`
	CanBeUpgraded           *bool           `json:"can_be_upgraded,omitempty"`
	Text                    *string         `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               *bool           `json:"is_private,omitempty"`
}

type UniqueGiftInfo struct {
	Gift              UniqueGift `json:"gift"`
	Origin            string     `json:"origin"`
	OwnedGiftID       *string    `json:"owned_gift_id,omitempty"`
	TransferStarCount *int       `json:"transfer_star_count,omitempty"`
}

type OwnedGift any

type OwnedGiftRegular struct {
	Type                    string          `json:"type"`
	Gift                    Gift            `json:"gift"`
	OwnedGiftID             *string         `json:"owned_gift_id,omitempty"`
	SenderUser              *User           `json:"sender_user,omitempty"`
	SendDate                *int            `json:"send_date,omitempty"`
	Text                    *string         `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               *bool           `json:"is_private,omitempty"`
	IsSaved                 *bool           `json:"is_saved,omitempty"`
	CanBeUpgraded           *bool           `json:"can_be_upgraded,omitempty"`
	WasRefunded             *bool           `json:"was_refunded,omitempty"`
	ConvertStarCount        *int            `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount *int            `json:"prepaid_upgrade_star_count,omitempty"`
}

type OwnedGiftUnique struct {
	Type              string     `json:"type"`
	Gift              UniqueGift `json:"gift"`
	OwnedGiftID       *string    `json:"owned_gift_id,omitempty"`
	SenderUser        *User      `json:"sender_user,omitempty"`
	SendDate          *int       `json:"send_date,omitempty"`
	IsSaved           *bool      `json:"is_saved,omitempty"`
	CanBeTransferred  *bool      `json:"can_be_transferred,omitempty"`
	TransferStarCount *int       `json:"transfer_star_count,omitempty"`
}

type OwnedGifts struct {
	TotalCount int         `json:"total_count"`
	Gifts      []OwnedGift `json:"gifts"`
	NextOffset *string     `json:"next_offset,omitempty"`
}

type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
}

type StarAmount struct {
	Amount         int  `json:"amount"`
	NanostarAmount *int `json:"nanostar_amount,omitempty"`
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
	UserID int64  `json:"user_id"`
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
	GiveawayMessageID int64  `json:"giveaway_message_id"`
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
// STICKER https://core.telegram.org/bots/api#stickers

type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            *string       `json:"emoji,omitempty"`
	SetName          *string       `json:"set_name,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID    *string       `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  *bool         `json:"needs_repainting,omitempty"`
	FileSize         *int64        `json:"file_size,omitempty"`
}

type StickerSet struct {
	Name        string     `json:"name"`
	Title       string     `json:"title"`
	StickerType string     `json:"sticker_type"`
	Stickers    []Sticker  `json:"stickers"`
	Thumbnail   *PhotoSize `json:"thumbnail,omitempty"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type InputSticker struct {
	Sticker      string        `json:"sticker"`
	Format       string        `json:"format"`
	EmojiList    []string      `json:"emoji_list"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	Keywords     []string      `json:"keywords,omitempty"`
}

//-------------------------------------------------------------
// INLINE TYPES https://core.telegram.org/bots/api#inline-mode

type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType *string   `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app,omitempty"`
	StartParameter *string     `json:"start_parameter,omitempty"`
}

type InlineQueryResult any

type InlineQueryResultArticle struct {
	Type                string                `json:"type"` // "article"
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 *string               `json:"url,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}
type InlineQueryResultPhoto struct {
	Type                  string                `json:"type"` // "photo"
	ID                    string                `json:"id"`
	PhotoURL              string                `json:"photo_url"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	PhotoWidth            *int                  `json:"photo_width,omitempty"`
	PhotoHeight           *int                  `json:"photo_height,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultGif struct {
	Type                  string                `json:"type"` // "gif"
	ID                    string                `json:"id"`
	GifURL                string                `json:"gif_url"`
	GifWidth              *int                  `json:"gif_width,omitempty"`
	GifHeight             *int                  `json:"gif_height,omitempty"`
	GifDuration           *int                  `json:"gif_duration,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     *string               `json:"thumbnail_mime_type,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultMpeg4Gif struct {
	Type                  string                `json:"type"` // "mpeg4_gif"
	ID                    string                `json:"id"`
	Mpeg4URL              string                `json:"mpeg4_url"`
	Mpeg4Width            *int                  `json:"mpeg4_width,omitempty"`
	Mpeg4Height           *int                  `json:"mpeg4_height,omitempty"`
	Mpeg4Duration         *int                  `json:"mpeg4_duration,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     *string               `json:"thumbnail_mime_type,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultVideo struct {
	Type                  string                `json:"type"` // "video"
	ID                    string                `json:"id"`
	VideoURL              string                `json:"video_url"`
	MimeType              string                `json:"mime_type"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	Title                 string                `json:"title"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	VideoWidth            *int                  `json:"video_width,omitempty"`
	VideoHeight           *int                  `json:"video_height,omitempty"`
	VideoDuration         *int                  `json:"video_duration,omitempty"`
	Description           *string               `json:"description,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultAudio struct {
	Type                string                `json:"type"` // "audio"
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       *int                  `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultVoice struct {
	Type                string                `json:"type"` // "voice"
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	VoiceDuration       *int                  `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultDocument struct {
	Type                string                `json:"type"` // "document"
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultLocation struct {
	Type                 string                `json:"type"` // "location"
	ID                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   *float64              `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int                  `json:"live_period,omitempty"`
	Heading              *int                  `json:"heading,omitempty"`
	ProximityAlertRadius *int                  `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbnailURL         *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      *int                  `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultVenue struct {
	Type                string                `json:"type"` // "venue"
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        *string               `json:"foursquare_id,omitempty"`
	FoursquareType      *string               `json:"foursquare_type,omitempty"`
	GooglePlaceID       *string               `json:"google_place_id,omitempty"`
	GooglePlaceType     *string               `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultContact struct {
	Type                string                `json:"type"` // "contact"
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            *string               `json:"last_name,omitempty"`
	Vcard               *string               `json:"vcard,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultGame struct {
	Type          string                `json:"type"` // "game"
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedPhoto struct {
	Type                  string                `json:"type"` // "photo"
	ID                    string                `json:"id"`
	PhotoFileID           string                `json:"photo_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedGif struct {
	Type                  string                `json:"type"` // "gif"
	ID                    string                `json:"id"`
	GifFileID             string                `json:"gif_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                  string                `json:"type"` // "mpeg4_gif"
	ID                    string                `json:"id"`
	Mpeg4FileID           string                `json:"mpeg4_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"` // "sticker"
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"` // "document"
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVideo struct {
	Type                  string                `json:"type"` // "video"
	ID                    string                `json:"id"`
	VideoFileID           string                `json:"video_file_id"`
	Title                 string                `json:"title"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"` // "voice"
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"` // "audio"
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

type InputMessageContent any

type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          *string             `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

type InputLocationMessageContent struct {
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int     `json:"live_period,omitempty"`
	Heading              *int     `json:"heading,omitempty"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"`
}

type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

type InputContactMessageContent struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	Vcard       *string `json:"vcard,omitempty"`
}

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             *string        `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"` // Assumes LabeledPrice is defined
	MaxTipAmount              *int           `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              *string        `json:"provider_data,omitempty"`
	PhotoURL                  *string        `json:"photo_url,omitempty"`
	PhotoSize                 *int           `json:"photo_size,omitempty"`
	PhotoWidth                *int           `json:"photo_width,omitempty"`
	PhotoHeight               *int           `json:"photo_height,omitempty"`
	NeedName                  *bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool          `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool          `json:"need_email,omitempty"`
	NeedShippingAddress       *bool          `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool          `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool          `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool          `json:"is_flexible,omitempty"`
}

type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`               // Assumes User is defined
	Location        *Location `json:"location,omitempty"` // Assumes Location is defined
	InlineMessageID *string   `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

type SentWebAppMessage struct {
	InlineMessageID *string `json:"inline_message_id,omitempty"`
}

type PreparedInlineMessage struct {
	ID             string `json:"id"`
	ExpirationDate int    `json:"expiration_date"`
}

//-------------------------------------------------------------
//PAYMENTS TYPES https://core.telegram.org/bots/api#payments

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type OrderInfo struct {
	Name            *string          `json:"name,omitempty"`
	PhoneNumber     *string          `json:"phone_number,omitempty"`
	Email           *string          `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`
	TotalAmount                int        `json:"total_amount"`
	InvoicePayload             string     `json:"invoice_payload"`
	SubscriptionExpirationDate *int64     `json:"subscription_expiration_date,omitempty"`
	IsRecurring                *bool      `json:"is_recurring,omitempty"`
	IsFirstRecurring           *bool      `json:"is_first_recurring,omitempty"`
	ShippingOptionID           *string    `json:"shipping_option_id,omitempty"`
	OrderInfo                  *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID    string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID    string     `json:"provider_payment_charge_id"`
}

type RefundedPayment struct {
	Currency                string  `json:"currency"`
	TotalAmount             int     `json:"total_amount"`
	InvoicePayload          string  `json:"invoice_payload"`
	TelegramPaymentChargeID string  `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID *string `json:"provider_payment_charge_id,omitempty"`
}

type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             User       `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID *string    `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

type PaidMediaPurchased struct {
	From             User   `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

type RevenueWithdrawalState any

type RevenueWithdrawalStatePending struct {
	Type string `json:"type"`
}

type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"`
	Date int64  `json:"date"`
	URL  string `json:"url"`
}

type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"`
}

type AffiliateInfo struct {
	AffiliateUser      *User `json:"affiliate_user,omitempty"`
	AffiliateChat      *Chat `json:"affiliate_chat,omitempty"`
	CommissionPerMille int   `json:"commission_per_mille"`
	Amount             int   `json:"amount"`
	NanostarAmount     *int  `json:"nanostar_amount,omitempty"`
}

type TransactionPartner any

type TransactionPartnerUser struct {
	Type                        string         `json:"type"`
	TransactionType             string         `json:"transaction_type"`
	User                        User           `json:"user"`
	Affiliate                   *AffiliateInfo `json:"affiliate,omitempty"`
	InvoicePayload              *string        `json:"invoice_payload,omitempty"`
	SubscriptionPeriod          *int           `json:"subscription_period,omitempty"`
	PaidMedia                   []PaidMedia    `json:"paid_media,omitempty"`
	PaidMediaPayload            *string        `json:"paid_media_payload,omitempty"`
	Gift                        *Gift          `json:"gift,omitempty"`
	PremiumSubscriptionDuration *int           `json:"premium_subscription_duration,omitempty"`
}

type TransactionPartnerChat struct {
	Type string `json:"type"`
	Chat Chat   `json:"chat"`
	Gift *Gift  `json:"gift,omitempty"`
}

type TransactionPartnerAffiliateProgram struct {
	Type               string `json:"type"`
	SponsorUser        *User  `json:"sponsor_user,omitempty"`
	CommissionPerMille int    `json:"commission_per_mille"`
}

type TransactionPartnerFragment struct {
	Type            string                 `json:"type"`
	WithdrawalState RevenueWithdrawalState `json:"withdrawal_state,omitempty"`
}

type TransactionPartnerTelegramAds struct {
	Type string `json:"type"`
}

type TransactionPartnerTelegramApi struct {
	Type         string `json:"type"`
	RequestCount int    `json:"request_count"`
}

type TransactionPartnerOther struct {
	Type string `json:"type"`
}

type StarTransaction struct {
	ID             string             `json:"id"`
	Amount         int                `json:"amount"`
	NanostarAmount *int               `json:"nanostar_amount,omitempty"`
	Date           int64              `json:"date"`
	Source         TransactionPartner `json:"source,omitempty"`
	Receiver       TransactionPartner `json:"receiver,omitempty"`
}

type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

//-------------------------------------------------------------
//TELEGRAM PASSPORT TYPES https://core.telegram.org/bots/api#telegram-passport

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        *string        `json:"data,omitempty"`
	PhoneNumber *string        `json:"phone_number,omitempty"`
	Email       *string        `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportElementError any

type PassportElementErrorDataField struct {
	Source    string `json:"source"`
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorSelfie struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

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
