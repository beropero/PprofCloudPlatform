import req from '@/utils/request'
import type { 
    CreateProjectReq,
    GetProjectByPageUserReq
} from './types';
import { getUser } from '@/stores/user';

export function GetProjectByPageUser(getProjectByPageUserReq: GetProjectByPageUserReq){
    return req({
        method: 'post',
        url: '/project/getprojectbypageuser',
        data: getProjectByPageUserReq,
        headers: {
            Authorization: getUser().token
        }
    });
}

export function CreateProject(CreateProjectreq: CreateProjectReq){
    return req({
        method: 'post',
        url: '/project/createproject',
        data: CreateProjectreq,
        headers: {
            Authorization: getUser().token
        }
    });
}
