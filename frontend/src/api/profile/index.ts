import req from '@/utils/request'
import type { 
    GetProfileByPageUserReq
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