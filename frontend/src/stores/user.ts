//保存用户jwt

export interface User {
    email: string;
    token: string;
}

// 保存用户信息
export const saveUser = (user: User) => {
    saveUserInfo(user);
}

export const saveUserInfo = (user: User) => {
    localStorage.setItem('userInfo', JSON.stringify(user));
}

// 获取用户信息
export const getUser = () => {
    return JSON.parse(localStorage.getItem('userInfo') || '{}');
}
