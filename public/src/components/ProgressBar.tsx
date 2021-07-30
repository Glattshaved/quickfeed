import React from "react"
import { useAppState, useActions } from "../overmind"
import { Submission } from "../../proto/ag/ag_pb"
import { getPassedTestsCount } from "../Helpers"
import { json } from "overmind"


export const ProgressBar = (props: {courseID: number, assignmentIndex: number, submission?: Submission, type: string}) => {
    const state = useAppState()

    const submission = props.submission ? props.submission : state.submissions[props.courseID][props.assignmentIndex]
    const assignment = state.assignments[props.courseID][props.assignmentIndex]
    const passedTests =  getPassedTestsCount(json(submission).getScoresList())


    // Returns a thin line to be used for labs in the NavBar
    if(props.type === "navbar") {
        let percentage = 0
        let score = 0
        if (submission) {
            percentage = 100 - submission.getScore()
            score = submission.getScore()
        }
        return (
            <div style={{ 
                position: "absolute", 
                borderBottom: "2px solid green", 
                bottom: 0, 
                left: 0, 
                right: `${percentage}%`, 
                borderColor: `${score >= assignment.getScorelimit() ? "green" : "yellow"}`, 
                opacity: 0.3 }}>
            </div>
        )
    }

    // Returns a regular size progress bar to be used for labs
    if(props.type === "lab") {
        let color: string
        switch (submission.getStatus()) {
            case Submission.Status.NONE:
                if(submission.getScore()>=state.assignments[props.courseID][props.assignmentIndex].getScorelimit()){
                    color = "bg-primary"
                }
                else{
                    color = "bg-secondary"
                }
                break
            case Submission.Status.APPROVED:
                color = "bg-success"
                break
            case Submission.Status.REJECTED:
                color = "bg-danger"
                break
            case Submission.Status.REVISION:
                color = "bg-warning text-dark"
                break
        }

        return (
            <div className="progress">
                <div 
                    className={"progress-bar "+color} 
                    role="progressbar" 
                    style={{width: props.submission?.getScore() + "%", transitionDelay: "0.5s"}} 
                    aria-valuenow={submission.getScore()} 
                    aria-valuemin={0} 
                    aria-valuemax={100}>
                        {props.submission?.getScore()}% ({passedTests})
                </div>
            </div>
        )
    }
    return (
        <div>
            
        </div>
    )
}
