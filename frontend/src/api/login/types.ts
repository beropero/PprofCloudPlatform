export interface LoginReq {
    email: string;
    password: string;
}

export interface RegisterReq {
    code: string;
    email: string;
    password: string;
}

export interface SendEmailReq {
    email: string;
}
