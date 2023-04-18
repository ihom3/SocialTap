import { Social } from "./Social"

export interface SocialReq {
    name: string,
    status: boolean,
    url: string
}


export interface UserRequest {
    active: boolean,
    unregistered: boolean,
    user: {
        bio_text: string,
        first_name: String,
        last_name: String, 
        sticker_code: String,
        profile_picture: string,
        social_list: {
            facebook?: SocialReq,
            snapchat?: SocialReq,
            instagram?: SocialReq
        }
    }
}