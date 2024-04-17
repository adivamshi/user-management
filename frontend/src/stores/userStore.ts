import {defineStore} from 'pinia'
import type {User} from "@/utils/types";

export type UserMap = {[id: number]: User}
// Not using pinia
export const useUsersStore = defineStore('Users',{
    state: () =>
        ({
            users: {} as UserMap,
            count: 0,
        }),
    actions:{
        setUsers(newUsers: User[], count: number = 0){
            const newUsersMap = newUsers.reduce((userObject: UserMap, user) => {
                userObject[user.id] = user
                return userObject;
            }, {} as UserMap);
            this.$state ={
                users:newUsersMap,
                count: count
            };
        },
        createUser(user: User){
            if (!user) return;
            this.users[user.id as number] = user;
            this.count++
        },
        updateUser(id:number, user: User){
            if (!id || !user) return;

            this.users[id] = user;
        },
        deleteUser(id:number){
            delete this.users[id];
            this.count--
        },
        usersList():User[]{
            return Object.values(this.users)
        },
        resetState() {
            this.$state = {
                users: {} as UserMap,
                count: 0
            };
        },

    }
})
