import React, { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import { useOvermind } from "../overmind"

/* 
    This component generates links for all labs for the sidebar. 
    Labs show depending on the state.activeCourse, and only labs for one course will be displayed at a time.
    Additionally, the lab that is selected from the list will be tagged as 'active', giving a visual que that it is the selected lab.

    // TODO Currently there is functionality implemented that displays a small 'progress bar', given by the submission score, at the bottom om each lab element. 
    the intention for this is to give the user a quick way to assess progress in a course. This might be overkill.
**/
const NavBarLabs = () => {
    const {state} = useOvermind()

    const [active, setActive] = useState(-1)
    let course = -1
    useEffect(() => {
        // If the active course changes, this prevents the previously selected lab to be active for the incoming course
        if (course !== state.activeCourse) {
            setActive(-1)
            course = state.activeCourse
        }
    }, [state.activeCourse])

    const labs = (): JSX.Element[] => { 
        
        if(state.submissions !== undefined && state.activeCourse > 0) {
            let links = state.submissions[state.activeCourse]?.map((submission, index) => {
                let assignment = state.assignments[state.activeCourse][index]
                let rand = Math.random()
                let percentage = 100 - (submission.getScore() - rand * 100)
                return (
                    <li style={{position: "relative"}} className={active === index && state.activeCourse === assignment.getCourseid() ? "active" : ""} key={submission.getAssignmentid()} onClick={() => {setActive(index)}}>
                        <div id="icon">
                            <i className={assignment.getIsgrouplab() ? "fa fa-users" : "fa fa-user"} title={assignment.getIsgrouplab() ? "Group Lab" : "Individual Lab"}></i>
                        </div>
                        <div id="title">
                            <Link to={`/course/${state.activeCourse}/${submission.getAssignmentid()}`}>{assignment.getName()}</Link>
                        </div> 
                        
                        {/** The following code adds a "progress" bar below the lab in the sidebar to indicate how many % done a user is with a lab where the percentage is (100 - submission.getScore()).
                         * if the score is above the assignment score limit, the bar turns green, else it is yellow. //TODO Should look into de-spaghettifying this */ }
                        <div style={{position: "absolute", borderBottom: "1px solid green", bottom: 0, left: 0, right: `${percentage}%`, borderColor: `${ submission.getScore() - rand * 100 >= assignment.getScorelimit() ? "green" : "yellow"}`, opacity:0.3 }}></div>
                    </li>
                )
            })
        return links
        }
        return []

    }
    return (
        <React.Fragment>
            {labs()}
        </React.Fragment>
    )
}

export default NavBarLabs