import React from "react"
import { Dispatch, SetStateAction } from "react"
import { GradingBenchmark, GradingCriterion } from "../../../proto/ag/ag_pb"
import { useActions, useAppState } from "../../overmind"

type GradeCommentProps = {
    grade: GradingBenchmark.AsObject | GradingCriterion.AsObject,
    editing: boolean,
    setEditing: Dispatch<SetStateAction<boolean>>
}

const GradeComment = ({ grade, editing, setEditing }: GradeCommentProps): JSX.Element => {
    const actions = useActions()
    const state = useAppState()

    /* Don't allow grading if user is not a teacher or editing is false */
    if (!state.isTeacher || !editing) {
        return <></>
    }

    /* Currently only triggers when user clicks outside of the textarea */
    const handleChange = (event: React.FormEvent<HTMLInputElement>) => {
        const { value } = event.currentTarget
        setEditing(false)
        // Exit early if the value is unchanged
        if (value === grade.comment) {
            return
        }
        actions.review.updateComment({ grade: grade, comment: value })
    }

    return (
        <tr>
            <th colSpan={3}>
                <input autoFocus onBlur={(e) => handleChange(e)} defaultValue={grade.comment} className="form-control" type="text" />
            </th>
        </tr>
    )

}

export default GradeComment
