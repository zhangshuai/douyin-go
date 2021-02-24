package conf

const Version = "0.0.1"

const (
	CONTENT_TYPE_JSON                 = "application/json"
	CONTENT_TYPE_FORM                 = "application/x-www-form-urlencoded"
	CONTENT_TYPE_OCTET                = "application/octet-stream"
	API_HOST                          = "open.douyin.com"
	API_HTTP_SCHEME                   = "https://"
	API_OAUTH_CONNECT                 = "/platform/oauth/connect/"
	API_OAUTH_ACCESS_TOKEN            = "/oauth/access_token/"
	API_OAUTH_REFRESH_TOKEN           = "/oauth/refresh_token/"
	API_OAUTH_RENEW_REFRESH_TOKEN     = "/oauth/renew_refresh_token/"
	API_OAUTH_USERINFO                = "/oauth/userinfo/"
	API_OAUTH_CLIENT_ACCESS_TOKEN     = "/oauth/client_token/"
	API_VIDEO_LIST                    = "/video/list/"
	API_VIDEO_UPLOAD                  = "/video/upload/"
	API_VIDEO_UPLOAD_PART_INIT        = "/video/part/init/"
	API_VIDEO_UPLOAD_PART_UPLOAD      = "/video/part/upload/"
	API_VIDEO_UPLOAD_PART_COMPLETE    = "/video/part/complete/"
	API_VIDEO_CREATE                  = "/video/create/"
	API_VIDEO_DELETE                  = "/video/delete/"
	API_VIDEO_DATA                    = "/video/data/"
	API_IMAGE_UPLOAD                  = "/image/upload/"
	API_IMAGE_CREATE                  = "/image/create/"
	API_AWEME_SHARE                   = "/share-id/"
	API_DATA_EXTERNAL_USER_ITEM       = "/data/external/user/item/"
	API_DATA_EXTERNAL_USER_FANS       = "/data/external/user/fans/"
	API_DATA_EXTERNAL_USER_LIKE       = "/data/external/user/like/"
	API_DATA_EXTERNAL_USER_COMMENT    = "/data/external/user/comment/"
	API_DATA_EXTERNAL_USER_SHARE      = "/data/external/user/share/"
	API_DATA_EXTERNAL_USER_PROFILE    = "/data/external/user/profile/"
	API_DATA_EXTERNAL_ITEM_BASE       = "/data/external/item/base/"
	API_DATA_EXTERNAL_ITEM_LIKE       = "/data/external/item/like/"
	API_DATA_EXTERNAL_ITEM_COMMENT    = "/data/external/item/comment/"
	API_DATA_EXTERNAL_ITEM_PLAY       = "/data/external/item/play/"
	API_DATA_EXTERNAL_ITEM_SHARE      = "/data/external/item/share/"
	API_HOT_SEARCH_SENTENCES          = "/hotsearch/sentences/"
	API_HOT_SEARCH_TRENDING_SENTENCES = "/hotsearch/trending/sentences/"
	API_HOT_SEARCH_VIDEOS             = "/hotsearch/videos/"
	API_FANS_LIST                     = "/fans/list/"
	API_FOLLOWING_LIST                = "/following/list/"
	API_ITEM_COMMENT_LIST             = "/item/comment/list/"
	API_ITEM_COMMENT_REPLY_LIST       = "/item/comment/reply/list/"
	API_ITEM_COMMENT_REPLY            = "/item/comment/reply/"
	API_EVENT_STATUS_LIST             = "/event/status/list/"
	API_EVENT_STATUS_UPDATE           = "/event/status/update/"
	API_STAR_HOT_LIST                 = "/star/hot_list/"
	API_STAR_AUTHOR_SCORE_V2          = "/star/author_score_v2/"
	API_STAR_AUTHOR_SCORE             = "/star/author_score/"
	API_JS_TICKET                     = "/js/getticket/"
)