package conf

const (
	// API_HOST OpenAPI HOST
	API_HOST = "open.douyin.com"

	// API_HTTP_SCHEME 协议
	API_HTTP_SCHEME = "https://"

	// API_OAUTH_CONNECT 生成授权链接
	API_OAUTH_CONNECT = "/platform/oauth/connect/"

	// API_OAUTH_ACCESS_TOKEN 获取access_token
	API_OAUTH_ACCESS_TOKEN = "/oauth/access_token/"

	// API_OAUTH_REFRESH_TOKEN 刷新access_token
	API_OAUTH_REFRESH_TOKEN = "/oauth/refresh_token/"

	// API_OAUTH_RENEW_REFRESH_TOKEN 刷新refresh_token
	API_OAUTH_RENEW_REFRESH_TOKEN = "/oauth/renew_refresh_token/"

	// API_OAUTH_USERINFO 获取用户信息
	API_OAUTH_USERINFO = "/oauth/userinfo/"

	// API_OAUTH_CLIENT_ACCESS_TOKEN 生成client_token
	API_OAUTH_CLIENT_ACCESS_TOKEN = "/oauth/client_token/"

	// API_VIDEO_LIST 查询授权账号视频数据
	API_VIDEO_LIST = "/video/list/"

	// API_VIDEO_UPLOAD 上传视频到文件服务器
	API_VIDEO_UPLOAD = "/video/upload/"

	// API_VIDEO_UPLOAD_PART_INIT 初始化分片上传
	API_VIDEO_UPLOAD_PART_INIT = "/video/part/init/"

	// API_VIDEO_UPLOAD_PART_UPLOAD 上传视频分片到文件服务器
	API_VIDEO_UPLOAD_PART_UPLOAD = "/video/part/upload/"

	// API_VIDEO_UPLOAD_PART_COMPLETE 完成上传视频
	API_VIDEO_UPLOAD_PART_COMPLETE = "/video/part/complete/"

	// API_VIDEO_CREATE 创建抖音视频
	API_VIDEO_CREATE = "/video/create/"

	// API_VIDEO_DELETE  删除授权用户发布的视频
	API_VIDEO_DELETE = "/video/delete/"

	// API_VIDEO_DATA 查询指定视频数据
	API_VIDEO_DATA = "/video/data/"

	// API_VIDEO_SEARCH 关键词视频搜索
	API_VIDEO_SEARCH = "/video/search/"

	// API_VIDEO_SEARCH_COMMENT_LIST 关键词视频评论列表
	API_VIDEO_SEARCH_COMMENT_LIST = "/video/search/comment/list/"

	// API_VIDEO_SEARCH_COMMENT_REPLY_LIST 关键词视频评论回复列表
	API_VIDEO_SEARCH_COMMENT_REPLY_LIST = "/video/search/comment/reply/list/"

	// API_VIDEO_SEARCH_COMMENT_REPLY 关键词视频评论回复
	API_VIDEO_SEARCH_COMMENT_REPLY = "/video/search/comment/reply/"

	// API_IMAGE_UPLOAD 上传图片到文件服务器
	API_IMAGE_UPLOAD = "/image/upload/"

	// API_IMAGE_CREATE 发布图片
	API_IMAGE_CREATE = "/image/create/"

	// API_AWEME_SHARE 获取share-id
	API_AWEME_SHARE = "/share-id/"

	// API_DATA_EXTERNAL_USER_ITEM 获取用户视频情况
	API_DATA_EXTERNAL_USER_ITEM = "/data/external/user/item/"

	// API_DATA_EXTERNAL_USER_FANS 获取用户粉丝数
	API_DATA_EXTERNAL_USER_FANS = "/data/external/user/fans/"

	// API_DATA_EXTERNAL_USER_LIKE 获取用户点赞数
	API_DATA_EXTERNAL_USER_LIKE = "/data/external/user/like/"

	// API_DATA_EXTERNAL_USER_COMMENT 获取用户评论数
	API_DATA_EXTERNAL_USER_COMMENT = "/data/external/user/comment/"

	// API_DATA_EXTERNAL_USER_SHARE 获取用户分享数
	API_DATA_EXTERNAL_USER_SHARE = "/data/external/user/share/"

	// API_DATA_EXTERNAL_USER_PROFILE 获取用户主页访问数
	API_DATA_EXTERNAL_USER_PROFILE = "/data/external/user/profile/"

	// API_DATA_EXTERNAL_ITEM_BASE 获取视频基础数据
	API_DATA_EXTERNAL_ITEM_BASE = "/data/external/item/base/"

	// API_DATA_EXTERNAL_ITEM_LIKE 获取视频点赞数据
	API_DATA_EXTERNAL_ITEM_LIKE = "/data/external/item/like/"

	// API_DATA_EXTERNAL_ITEM_COMMENT 获取视频评论数据
	API_DATA_EXTERNAL_ITEM_COMMENT = "/data/external/item/comment/"

	// API_DATA_EXTERNAL_ITEM_PLAY 获取视频播放数据
	API_DATA_EXTERNAL_ITEM_PLAY = "/data/external/item/play/"

	// API_DATA_EXTERNAL_ITEM_SHARE 获取视频分享数据
	API_DATA_EXTERNAL_ITEM_SHARE = "/data/external/item/share/"

	// API_HOT_SEARCH_SENTENCES 获取实时热点词
	API_HOT_SEARCH_SENTENCES = "/hotsearch/sentences/"

	// API_HOT_SEARCH_TRENDING_SENTENCES 获取上升词
	API_HOT_SEARCH_TRENDING_SENTENCES = "/hotsearch/trending/sentences/"

	// API_HOT_SEARCH_VIDEOS 获取热点词聚合的视频
	API_HOT_SEARCH_VIDEOS = "/hotsearch/videos/"

	// API_FANS_LIST 获取粉丝列表
	API_FANS_LIST = "/fans/list/"

	// API_FANS_DATA 获取用户粉丝数据
	API_FANS_DATA = "/fans/data/"

	// API_FANS_CHECK 获取粉丝判断
	API_FANS_CHECK = "/fans/check/"

	// API_DATA_EXTERNAL_FANS_SOURCE 获取用户粉丝来源分布
	API_DATA_EXTERNAL_FANS_SOURCE = "/data/extern/fans/source/"

	// API_DATA_EXTERNAL_FANS_FAVOURITE 获取用户粉丝喜好
	API_DATA_EXTERNAL_FANS_FAVOURITE = "/data/extern/fans/favourite/"

	// API_DATA_EXTERNAL_FANS_COMMENT 获取用户粉丝热评
	API_DATA_EXTERNAL_FANS_COMMENT = "/data/extern/fans/comment/"

	// API_FOLLOWING_LIST 获取关注列表
	API_FOLLOWING_LIST = "/following/list/"

	// API_ITEM_COMMENT_LIST 获取评论列表
	API_ITEM_COMMENT_LIST = "/item/comment/list/"

	// API_ITEM_COMMENT_REPLY_LIST 获取评论回复列表
	API_ITEM_COMMENT_REPLY_LIST = "/item/comment/reply/list/"

	// API_ITEM_COMMENT_REPLY 回复视频评论
	API_ITEM_COMMENT_REPLY = "/item/comment/reply/"

	// API_EVENT_STATUS_LIST 获取事件订阅状态
	API_EVENT_STATUS_LIST = "/event/status/list/"

	// API_EVENT_STATUS_UPDATE 更新应用推送事件订阅状态
	API_EVENT_STATUS_UPDATE = "/event/status/update/"

	// API_STAR_HOT_LIST 获取抖音星图达人热榜
	API_STAR_HOT_LIST = "/star/hot_list/"

	// API_STAR_AUTHOR_SCORE_V2 获取抖音星图达人指数数据V2
	API_STAR_AUTHOR_SCORE_V2 = "/star/author_score_v2/"

	// API_STAR_AUTHOR_SCORE 获取抖音星图达人指数
	API_STAR_AUTHOR_SCORE = "/star/author_score/"

	// API_JS_TICKET 获取jsapi_ticket
	API_JS_TICKET = "/js/getticket/"

	// API_DISCOVERY_ENT_RANK_ITEM 获取抖音电影榜、抖音电视剧榜、抖音综艺榜
	API_DISCOVERY_ENT_RANK_ITEM = "/discovery/ent/rank/item/"

	// API_DISCOVERY_ENT_RANK_VERSION 获取抖音影视综榜单版本
	API_DISCOVERY_ENT_RANK_VERSION = "/discovery/ent/rank/version/"

	// API_DATA_EXTERNAL_BILLBOARD_GAME_CONSOLE 获取单机主机榜数据
	API_DATA_EXTERNAL_BILLBOARD_GAME_CONSOLE = "/data/extern/billboard/game/console/"

	// API_DATA_EXTERNAL_BILLBOARD_GAME_INFO 获取游戏资讯数据
	API_DATA_EXTERNAL_BILLBOARD_GAME_INFO = "/data/extern/billboard/game/inf/"

	// API_DATA_EXTERNAL_BILLBOARD_DRAME_OVERALL 获取剧情总榜数据
	API_DATA_EXTERNAL_BILLBOARD_DRAME_OVERALL = "/data/extern/billboard/drama/overall/"

	// API_DATA_EXTERNAL_BILLBOARD_CAR_OVERALL 获取汽车总榜数据
	API_DATA_EXTERNAL_BILLBOARD_CAR_OVERALL = "/data/extern/billboard/car/overall/"

	// API_DATA_EXTERNAL_BILLBOARD_CAR_COMMENT 获取评车榜数据
	API_DATA_EXTERNAL_BILLBOARD_CAR_COMMENT = "/data/extern/billboard/car/comment/"

	// API_DATA_EXTERNAL_BILLBOARD_CAR_PLAY 获取玩车榜数据
	API_DATA_EXTERNAL_BILLBOARD_CAR_PLAY = "/data/extern/billboard/car/play/"

	// API_DATA_EXTERNAL_BILLBOARD_CAR_USE 获取用车榜数据
	API_DATA_EXTERNAL_BILLBOARD_CAR_USE = "/data/extern/billboard/car/use/"

	// API_DATA_EXTERNAL_BILLBOARD_CAR_DRIVER 获取驾考榜数据
	API_DATA_EXTERNAL_BILLBOARD_CAR_DRIVER = "/data/extern/billboard/car/driver/"

	// API_DATA_EXTERNAL_BILLBOARD_AMUSEMENT_OVERALL 获取搞笑总榜数据
	API_DATA_EXTERNAL_BILLBOARD_AMUSEMENT_OVERALL = "/data/extern/billboard/amusement/overall/"

	// API_DATA_EXTERNAL_BILLBOARD_AMUSEMENT_NEW 获取搞笑新势力数据
	API_DATA_EXTERNAL_BILLBOARD_AMUSEMENT_NEW = "/data/extern/billboard/amusement/new/"

	// API_DATA_EXTERNAL_BILLBOARD_COSPA_OVERALL 获取二次元总榜数据
	API_DATA_EXTERNAL_BILLBOARD_COSPA_OVERALL = "/data/extern/billboard/cospa/overall/"

	// API_DATA_EXTERNAL_BILLBOARD_COSPA_QING_MAN 获取轻漫榜数据
	API_DATA_EXTERNAL_BILLBOARD_COSPA_QING_MAN = "/data/extern/billboard/cospa/qing_man/"

	// API_DATA_EXTERNAL_BILLBOARD_COSPA_OUT_SHOT 获取出境拍摄榜数据
	API_DATA_EXTERNAL_BILLBOARD_COSPA_OUT_SHOT = "/data/extern/billboard/cospa/out_shot/"

	// API_DATA_EXTERNAL_BILLBOARD_COSPA_PAINTING 获取绘画榜数据
	API_DATA_EXTERNAL_BILLBOARD_COSPA_PAINTING = "/data/extern/billboard/cospa/painting/"

	// API_DATA_EXTERNAL_BILLBOARD_COSPA_VOICE_CONTROL 获取声控榜数据
	API_DATA_EXTERNAL_BILLBOARD_COSPA_VOICE_CONTROL = "/data/extern/billboard/cospa/voice_control/"

	// API_DATA_EXTERNAL_BILLBOARD_COSPA_BRAIN_CAVITY 获取脑洞榜数据
	API_DATA_EXTERNAL_BILLBOARD_COSPA_BRAIN_CAVITY = "/data/extern/billboard/cospa/brain_cavity/"

	// API_DATA_EXTERNAL_BILLBOARD_COSPA_NEW 获取二次元新势力榜数据
	API_DATA_EXTERNAL_BILLBOARD_COSPA_NEW = "/data/extern/billboard/cospa/new/"

	// API_DATA_EXTERNAL_BILLBOARD_FOOD_OVERALL 获取美食总榜数据
	API_DATA_EXTERNAL_BILLBOARD_FOOD_OVERALL = "/data/extern/billboard/food/overall/"

	// API_DATA_EXTERNAL_BILLBOARD_FOOD_NEW 获取美食新势力榜数据
	API_DATA_EXTERNAL_BILLBOARD_FOOD_NEW = "/data/extern/billboard/food/new/"

	// API_DATA_EXTERNAL_BILLBOARD_FOOD_TUTORIAL 获取美食教程榜数据
	API_DATA_EXTERNAL_BILLBOARD_FOOD_TUTORIAL = "/data/extern/billboard/food/tutorial/"

	// API_DATA_EXTERNAL_BILLBOARD_FOOD_SHOP 获取美食探店榜数据
	API_DATA_EXTERNAL_BILLBOARD_FOOD_SHOP = "/data/extern/billboard/food/shop/"

	// API_DATA_EXTERNAL_BILLBOARD_TRAVEL_OVERALL 获取旅游总榜数据
	API_DATA_EXTERNAL_BILLBOARD_TRAVEL_OVERALL = "/data/extern/billboard/travel/overall/"

	// API_DATA_EXTERNAL_BILLBOARD_TRAVEL_NEW 获取旅游新势力榜数据
	API_DATA_EXTERNAL_BILLBOARD_TRAVEL_NEW = "/data/extern/billboard/travel/new/"

	// API_DATA_EXTERNAL_BILLBOARD_STARS 获取娱乐明星榜数据
	API_DATA_EXTERNAL_BILLBOARD_STARS = "/data/extern/billboard/stars/"

	// API_DATA_EXTERNAL_BILLBOARD_SPORT_OVERALL 获取体育总榜数据
	API_DATA_EXTERNAL_BILLBOARD_SPORT_OVERALL = "/data/extern/billboard/sport/overall/"

	// API_DATA_EXTERNAL_BILLBOARD_SPORT_BASKETBALL 获取篮球榜数据
	API_DATA_EXTERNAL_BILLBOARD_SPORT_BASKETBALL = "/data/extern/billboard/sport/basketball/"

	// API_DATA_EXTERNAL_BILLBOARD_SPORT_SOCCER 获取足球榜数据
	API_DATA_EXTERNAL_BILLBOARD_SPORT_SOCCER = "/data/extern/billboard/sport/soccer/"

	// API_DATA_EXTERNAL_BILLBOARD_SPORT_COMPREHENSIVE 获取综合体育榜数据
	API_DATA_EXTERNAL_BILLBOARD_SPORT_COMPREHENSIVE = "/data/extern/billboard/sport/comprehensive/"

	// API_DATA_EXTERNAL_BILLBOARD_SPORT_FITNESS 获取运动健身榜数据
	API_DATA_EXTERNAL_BILLBOARD_SPORT_FITNESS = "/data/extern/billboard/sport/fitness/"

	// API_DATA_EXTERNAL_BILLBOARD_SPORT_OUTDOORS 获取户外运动榜数据
	API_DATA_EXTERNAL_BILLBOARD_SPORT_OUTDOORS = "/data/extern/billboard/sport/outdoors/"

	// API_DATA_EXTERNAL_BILLBOARD_SPORT_TABLE_TENNIES 获取台球榜数据(非网球数据)
	API_DATA_EXTERNAL_BILLBOARD_SPORT_TABLE_TENNIES = "/data/extern/billboard/sport/table_tennis/"

	// API_DATA_EXTERNAL_BILLBOARD_SPORT_CULTURE 获取运动文化榜数据
	API_DATA_EXTERNAL_BILLBOARD_SPORT_CULTURE = "/data/extern/billboard/sport/culture/"

	// API_DATA_EXTERNAL_BILLBOARD_TOPIC 获取话题榜数据
	API_DATA_EXTERNAL_BILLBOARD_TOPIC = "/data/extern/billboard/topic/"

	// API_DATA_EXTERNAL_BILLBOARD_PROP 获取话题榜数据
	API_DATA_EXTERNAL_BILLBOARD_PROP = "/data/extern/billboard/prop/"

	// API_DATA_EXTERNAL_BILLBOARD_HOT_VIDEO 获取热门视频数据
	API_DATA_EXTERNAL_BILLBOARD_HOT_VIDEO = "/data/extern/billboard/hot_video/"

	// API_DATA_EXTERNAL_BILLBOARD_LIVE 获取直播榜数据
	API_DATA_EXTERNAL_BILLBOARD_LIVE = "/data/extern/billboard/live/"

	// API_DATA_EXTERNAL_BILLBOARD_MUSIC_HOT 获取音乐热歌榜数据
	API_DATA_EXTERNAL_BILLBOARD_MUSIC_HOT = "/data/extern/billboard/music/hot/"

	// API_DATA_EXTERNAL_BILLBOARD_MUSIC_SOAR 获取音乐飙升榜数据
	API_DATA_EXTERNAL_BILLBOARD_MUSIC_SOAR = "/data/extern/billboard/music/soar/"

	// API_DATA_EXTERNAL_BILLBOARD_MUSIC_ORIGINAL 获取音乐原创榜数据
	API_DATA_EXTERNAL_BILLBOARD_MUSIC_ORIGINAL = "/data/extern/billboard/music/original/"

	// API_POI_SEARCH_KEYWORD 查询POI信息
	API_POI_SEARCH_KEYWORD = "/poi/search/keyword/"
)
