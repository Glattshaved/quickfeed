import React from "react"
import { GradingCriterion, Review } from "../../proto/ag/ag_pb"


interface submission {
    review: Review[]
}

const ReviewResult = ({review}: submission) => {
    const Result: Function = (): JSX.Element[] => {
        let b: JSX.Element[] = []

        review.forEach(r => {
            try {
                

        
            r.getBenchmarksList().map(benchmark => {
                b.push(
                <tr className="table-info">
                    <th colSpan={2}>{benchmark.getHeading()}</th>
                    <th>{benchmark.getComment()}</th>
                </tr>)
                benchmark.getCriteriaList().map(criteria => {
                    const passed = criteria.getGrade() == GradingCriterion.Grade.PASSED
                    const boxShadow = passed ? "0 0px 0 #000 inset, 5px 0 0 green inset" : "0 0px 0 #000 inset, 5px 0 0 red inset"
                    const icon = passed ? "fa fa-check" : "fa fa-exclamation-circle"
                    b.push(
                        <tr>
                            <th style={{boxShadow: boxShadow}}>{criteria.getDescription()} {criteria.getComment()}</th>
                            <th><i className={icon}></i></th>
                            <th>{criteria.getComment()}</th>
                        </tr>)
                })
            })
        
            }catch (error) {
                console.log(error)
            }
        })
        return b
    }

    return (
        <div className="container">
            <table className="table"> 
                <thead className="thead-dark">
                    {review.length > 0 ?
                    <tr className="table-primary">
                        <th>{review[0].getFeedback()}</th>
                        <th>{review[0].getScore()}%</th>
                        <th></th>
                    </tr>
                    : null
                    }  
                    <tr>
                        <th scope="col">Criteria</th>
                        <th scope="col">Status</th>
                        <th scope="col">Comment</th>
                    </tr>
                </thead>
                <Result />
            </table>    

        </div>
    )
}

export default ReviewResult