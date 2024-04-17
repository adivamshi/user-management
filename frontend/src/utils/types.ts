export interface User {
    id: number
    name: string
    email: string
    age:number
    created:number
    updated:number
}

export interface GetUserDataResponse {
    total:number
    users:User[]
}

export interface UpdateUserResponse {
    total:number
    updated:User
}
