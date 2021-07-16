import React, { useEffect } from "react"
import { useOvermind } from "../overmind"
import Group from "./group/Group"
import CreateGroup from "./group/CreateGroup"
import { getCourseID } from "../Helpers"
import { Enrollment } from "../../proto/ag/ag_pb"
import Groups from "./Groups"


export const GroupPage = () => {
    const {state, actions} = useOvermind()
    const courseID = getCourseID()

    useEffect(() => {
        actions.getGroupByUserAndCourse(courseID)
    })

    if (state.enrollmentsByCourseId[courseID].getStatus() == Enrollment.UserStatus.TEACHER) {
        return <Groups courseID={courseID}></Groups>
    }

    if (!state.userGroup[courseID]) {
        return <CreateGroup courseID={courseID} />    
    }
    return <Group courseID={courseID} />
    
}

export default GroupPage