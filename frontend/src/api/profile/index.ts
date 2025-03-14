import req from '@/utils/request'
import type { 
    GetProfileByPageUserReq,
    GetProfileContentReq
} from './types';
import { getUser } from '@/stores/user';

export function GetProfileByPageUser(getprofilebypageuser: GetProfileByPageUserReq){
    return req({
        method: 'post',
        url: '/profile/getprofilebypageuser',
        data: getprofilebypageuser,
        headers: {
            Authorization: getUser().token
        }
    });
}

export function GetProfileContent(getProfileContentReq: GetProfileContentReq){
    return req({
        method: 'get',
        url: '/profile/getprofilecontent',
        params: getProfileContentReq,
        headers: {
            Authorization: getUser().token
        }
    });
}