import React, { useEffect, useState } from "react"
import { RouteComponentProps } from "react-router"
import { Link } from "react-router-dom"
import { useOvermind } from "../overmind"

import { Enrollment, Repository } from "../proto/ag_pb"
import LandingPageLabTable from "./LandingPageLabTable"


interface MatchProps {
    id: string
}


const Course = (props: RouteComponentProps<MatchProps>) => {
    const { state, actions } = useOvermind()
    let courseID = Number(props.match.params.id)


    useEffect(() => {
        
    }, [props])

    if (state.courses){
        return (
        <div className="box">
            <h1>{state.enrollmentsByCourseId[courseID].getCourse()?.getName()}</h1>
            <div className="Links">
                <a href={state.repositories[courseID][Repository.Type.USER]}>User Repository</a>
                <a href={state.repositories[courseID][Repository.Type.GROUP]}>Group Repository</a>
                <a href={state.repositories[courseID][Repository.Type.COURSEINFO]}>Course Info</a>
                <a href={state.repositories[courseID][Repository.Type.ASSIGNMENTS]}>Assignments</a>
                <Link to={"/course/" + courseID + "/group"} >Group</Link>
            </div>
            
            <LandingPageLabTable courseID={courseID} />
            

        </div>)
    }
    return <h1>Loading</h1>
}

export default Course