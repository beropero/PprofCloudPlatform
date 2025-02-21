import req from '@/utils/request'
import type { 
    LoginReq, 
    RegisterReq 
} from './types';

// 登录接口
export function Login(loginReq: LoginReq){
    return req({
        method: 'post',
        url: '/login/login',
        data: loginReq,
    });
}


// 注册接口
export function Register(registerReq: RegisterReq){
    return req({
        method: 'post',
        url: '/login/register',
        data: registerReq,
    });
}

// 获取验证码接口
export function GetEmailCode(email: string){
    return req({
        method: 'post',
        url: '/login/getemailcode',
        params: {
            email
        }
    });
}
