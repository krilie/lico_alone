import store from 'storejs';

const UserToken = "UserKey";
let UserTokenStr = "";
const UserTraceId = "UserTraceId";

/**
 * @return {string}
 */
export function GetUserToken() {
    if (UserTokenStr === "")
        UserTokenStr = store.get(UserToken) ?? "";
    return UserTokenStr;
}

// customer id
export function GetCustomerTraceId() {
    let traceId = store.get(UserTraceId) ?? ""
    if (traceId === "") {
        traceId = generateUUID
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
    return token !== "";
}

/**
 * generateUUID 生成UUID
 * @returns {string} 返回字符串
 */
function generateUUID(){
    var d = new Date().getTime();
    var uuid = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = (d + Math.random()*16)%16 | 0;
        d = Math.floor(d/16);
        return (c=='x' ? r : (r&0x7|0x8)).toString(16);
    });
    return uuid;
}