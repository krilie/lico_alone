import store from 'storejs';

const UserToken = "UserKey";
let UserTokenStr = "";

/**
 * @return {string}
 */
export function GetUserToken() {
    if (UserTokenStr === "")
        UserTokenStr = store.get(UserToken) ?? "";
    return UserTokenStr;
}

export function SetUserToken(jwtToken) {
    UserTokenStr = jwtToken;
    return store.set(UserToken, jwtToken);
}

export function ClearToken() {
    UserTokenStr = "";
    return store.set(UserToken, "");
}

export function hasToken ()  {
    let token = GetUserToken();
    return token !== "";
}