import req from '@/utils/request'
import type { 
    CreateMicroserviceReq,
    GetMicroserviceByPageUserReq
} from './types';
import { getUser } from '@/stores/user';

export function GetMicroserviceByPageUser(getmicroservicebypageuser: GetMicroserviceByPageUserReq){
    return req({
        method: 'post',
        url: '/microservice/getmicroservicebypageuser',
        data: getmicroservicebypageuser,
        headers: {
            Authorization: getUser().token
        }
    });
}

export function CreateMicroservice(createmicroservicereq: CreateMicroserviceReq){
    return req({
        method: 'post',
        url: '/microservice/createmicroservice',
        data: createmicroservicereq,
        headers: {
            Authorization: getUser().token
        }
    });
}