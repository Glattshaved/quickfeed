import React from "react";
import { getFormattedDeadline } from "../Helpers";
import { useOvermind, useReaction } from "../overmind";

const LandingPageLabTable = () => {
    const { state } = useOvermind()
    
    //replace {} with a type of dictionary/record
    
    const makeTable = (): JSX.Element[] => {
        let table: JSX.Element[] = []
        for (const courseID in state.assignments) {
            state.assignments[courseID].map(assignment => {
                if(state.submissions[courseID]) {
                    let submission = state.submissions[courseID].find(submission => assignment.getId() === submission.getAssignmentid())
                
                if(submission){
                table.push(
                    <tr key = {assignment.getId()} className= {"clickable-row "}>
                        <td>{}</td>
                        <td>{assignment.getName()}</td>
                        <td>{submission.getScore()} / {assignment.getScorelimit()}</td>
                        <td>{getFormattedDeadline(assignment.getDeadline())}</td>
                        <td></td>
                        <td>{(assignment.getAutoapprove()==false && submission.getScore()>= assignment.getScorelimit()) ? "Awating approval":(assignment.getAutoapprove()==true && submission.getScore()>= assignment.getScorelimit())? "Approved(Auto approve)(shouldn't be in final version)":"Score not high enough"}</td>
                        <td>{assignment.getIsgrouplab() ? "Yes": "No"}</td>
                    </tr>
                )
                }
            }
            })
        }
        return table

    }
    
    return (
        <div>
            <table className="table" id="LandingPageTable">
                <thead>
                    <tr>
                        <th>Course</th>
                        <th>Assignment Name</th>
                        <th>Progress</th>
                        <th>Deadline</th>
                        <th>Time Left</th>
                        <th>Status</th>
                        <th>Grouplab</th>
                    </tr>
                </thead>
                <tbody>
                    {makeTable()}
                </tbody>
            </table>
        </div>
    )
}

export default LandingPageLabTable