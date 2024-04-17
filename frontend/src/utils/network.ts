import type {GetUserDataResponse, UpdateUserResponse, User} from "@/utils/types";

const backend_url: string = "http://localhost:8081";

export const GetUserData = (query: string = "", page: number, pageSize: number): Promise<GetUserDataResponse> =>
    new Promise<GetUserDataResponse>((resolve, reject) => {
        let url = `${backend_url}/api/v1/users/?page=${page}&pageSize=${pageSize}`;
        if (query.length > 0) {
            url = `${backend_url}/api/v1/users/search?page=${page}&pageSize=${pageSize}&query=${query}`;
        }
        fetch(url)
            .then(response => {
                if (response.ok) return response.json();
                reject(new Error('Network response was not ok.'));
            })
            .then(data => resolve(data))
            .catch(err => reject(err));
    });

export const CreateUser = (name: string, email: string, age: number): Promise<void> =>
    new Promise<void>((resolve, reject) => {
        fetch(`${backend_url}/api/v1/users/`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name, email, age }),
            redirect: "follow"
        })
            .then(response => {
                if (response.ok) return response.json();
                reject(new Error('Failed to create user'));
            })
            .then(data => resolve(data))
            .catch(error => reject(error));
    });

export const UpdateUser = (userId: number, name: string, email: string, age: number): Promise<UpdateUserResponse> =>
    new Promise<UpdateUserResponse>((resolve, reject) => {
        fetch(`${backend_url}/api/v1/users/${userId}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name, email, age }),
            redirect: "follow"
        })
            .then(response => {
                if (response.ok) return response.json();
                reject(new Error('Failed to update user'));
            })
            .then(data => resolve(data as UpdateUserResponse))
            .catch(error => reject(error));
    });

export const GetUserById = (userId: number): Promise<User> =>
    new Promise<any>((resolve, reject) => {
        fetch(`${backend_url}/api/v1/users/${userId}`)
            .then(response => {
                if (response.ok) return response.json();
                reject(new Error('Failed to fetch user'));
            })
            .then(data => resolve(data))
            .catch(error => reject(error));
    });

export const DeleteUser = (userId: number): Promise<void> =>
    new Promise<void>((resolve, reject) => {
        fetch(`${backend_url}/api/v1/users/${userId}`, {
            method: "DELETE",
            redirect: "follow"
        })
            .then(response => {
                if (response.ok) return response.json();
                reject(new Error('Failed to delete user'));
            })
            .then(data => resolve(data))
            .catch(error => reject(error));
    });
