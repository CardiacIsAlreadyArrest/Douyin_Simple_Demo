namespace go relation

struct RelationActionRequest {
    1: string token,              // 用户鉴权token
    2: i64 to_user_id,            // 对方用户id
    3: i32 action_type            // 1-关注，2-取消关注
}

struct RelationActionResponse {
    1: i32 status_code,           // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}

struct RelationFollowListRequest {
    1: i64 user_id,               // 用户id
    2: string token               // 用户鉴权token
}

struct RelationFollowListResponse {
    1: i32 status_code,           // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: list<User> user_list       // 用户信息列表
}

struct RelationFollowerListRequest {
    1: i64 user_id,               // 用户id
    2: string token               // 用户鉴权token
}

struct RelationFollowerListResponse {
    1: i32 status_code,           // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: list<User> user_list       // 用户列表
}

struct RelationFriendListRequest {
    1: i64 user_id,               // 用户id
    2: string token               // 用户鉴权token
}

struct RelationFriendListResponse {
    1: i32 status_code,           // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
    3: list<FriendUser> user_list // 用户列表
}

struct User {
    1: i64 id,                    // 用户id
    2: string name,               // 用户名称
    3: optional i64 follow_count,  // 关注总数
    4: optional i64 follower_count, // 粉丝总数
    5: bool is_follow,            // true-已关注，false-未关注
    6: optional string avatar,     // 用户头像
    7: optional string background_image, // 用户个人页顶部大图
    8: optional string signature,   // 个人简介
    9: optional i64 total_favorited, // 获赞数量
    10: optional i64 work_count,    // 作品数量
    11: optional i64 favorite_count  // 点赞数量
}

struct FriendUser {
    1: User base,
    2: string message,    // 和该好友的最新聊天消息
    3: i64 msgType                 // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

service RelationService {
    // 关注操作
    RelationActionResponse RelationAction(1: RelationActionRequest req)
    // 获取关注列表
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest req)
    // 获取粉丝列表
    RelationFollowerListResponse RelationFollowerList(1: RelationFollowerListRequest req)
    // 获取好友列表
	RelationFriendListResponse RelationFriendList(1:RelationFriendListRequest req)
}