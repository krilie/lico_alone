import store from 'storejs';

const UserToken = "UserKey";
let UserTokenStr = "";
const UserTraceId = "UserTraceId";

/**
 * @return {string}
 */
export function GetUserToken() {
    if (UserTokenStr === "") {
        UserTokenStr = store.get(UserToken) ?? "";
        if (UserTokenStr === undefined)
            UserTokenStr = "";
    }
    return UserTokenStr;
}

// customer id
export function GetCustomerTraceId() {
    let traceId = store.get(UserTraceId) ?? ""
    if (traceId === "") {
        traceId = generateUUID()
        store.set(UserTraceId, traceId)
    }
    return traceId;
}

export function SetUserToken(jwtToken) {
    UserTokenStr = jwtToken;
    return store.set(UserToken, jwtToken);
}

export function ClearToken() {
    UserTokenStr = "";
    return store.set(UserToken, "");
}

export function hasToken() {
    let token = GetUserToken();
    return token !== "" && token !== undefined;
}

/**
 * generateUUID 生成UUID
 * @returns {string} 返回字符串
 */
function generateUUID() {
    var s = [];
    var hexDigits = "0123456789abcdef";
    for (var i = 0; i < 36; i++) {
        s[i] = hexDigits.substr(Math.floor(Math.random() * 0x10), 1);
    }
    s[14] = "4";  // bits 12-15 of the time_hi_and_version field to 0010
    s[19] = hexDigits.substr((s[19] & 0x3) | 0x8, 1);  // bits 6-7 of the clock_seq_hi_and_reserved to 01
    s[8] = s[13] = s[18] = s[23] = "";    //"" 引号里面可以加任意字符，代表拼接的意思，如果不加就是 纯32位支付

    var uuid = s.join("");
    return uuid;
}