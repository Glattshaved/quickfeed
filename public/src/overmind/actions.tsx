import { Context, Action } from "overmind";
import { Courses, Course, User } from "../proto/ag_pb";
import { useEffects } from ".";


export const getUser: Action<void, Promise<boolean>> = ({state, effects}) => {
    return effects.api.getUser()
    .then((user) => {
        if (user.id == undefined) {
            return false
        }
        state.user = user;
        return true
    })
    
}
export const getUsers: Action<void> = ({state, effects}) => {
    state.users = []
    effects.api.getUsers(state).then(users => {
        users.forEach(user => {
            if (user.getStudentid() != "") {
                state.users.push(user)
            }
        });
    })
}

export const getCourses: Action<void> = ({state, effects}) => {
    state.courses = []
    effects.api.getCourses(state).then(courses => {
            courses.forEach(course => {
                console.log(course)
                state.courses.push(course)
            })
        });
}

export const setCurrentPage: Action<string> = ({state}, page) => {
    state.currentPage = page;
}