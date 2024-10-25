import React from "react"
import { Assignment, AssignmentSchema, Submission, SubmissionSchema } from "../../proto/qf/types_pb"
import ProgressBar, { Progress } from "../components/ProgressBar"
import { initializeOvermind } from "./TestHelpers"
import { Provider } from "overmind-react"
import { render } from "@testing-library/react"
import { create } from "@bufbuild/protobuf"

type ProgressBarTest = {
    desc: string,
    submission: Submission,
    assignment: Assignment,
    want: string
    assignmentIndex?: number
}


describe("ProgressBar", () => {
    React.useLayoutEffect = React.useEffect


    const progressBarTests: ProgressBarTest[] = [
        {
            desc: "100% Progress Bar",
            submission: create(SubmissionSchema,{
                score: 100,
            }),
            assignment: create(AssignmentSchema, { scoreLimit: 100 }),
            want: "100 %"
        },
        {
            desc: "0% Progress Bar",
            submission: create(SubmissionSchema, { score: 0 }),
            assignment: create(AssignmentSchema, { scoreLimit: 100 }),
            want: "0 %"
        },
        {
            desc: "50% Progress Bar",
            submission: create(SubmissionSchema, { score: 50 }),
            assignment: create(AssignmentSchema, { scoreLimit: 100 }),
            want: "50 %"
        },
        {
            desc: "50% Progress Bar, with 75% scorelimit",
            submission: create(SubmissionSchema, { score: 50 }),
            assignment: create(AssignmentSchema, { scoreLimit: 75 }),
            want: "50 %"
        },
        {
            desc: "75% Progress Bar, with 50% scorelimit",
            submission: create(SubmissionSchema, { score: 75 }),
            assignment: create(AssignmentSchema, { scoreLimit: 50 }),
            want: "75 %"
        },
        {
            desc: "75% Progress Bar, with 75% scorelimit",
            submission: create(SubmissionSchema, { score: 75 }),
            assignment: create(AssignmentSchema, { scoreLimit: 75 }),
            want: "75 %"
        },
        {
            desc: "Progress Bar without score",
            submission: create(SubmissionSchema),
            assignment: create(AssignmentSchema, { scoreLimit: 100 }),
            want: "0 %"
        },
        {
            desc: "Progress Bar without scorelimit",
            submission: create(SubmissionSchema, { score: 50 }),
            assignment: create(AssignmentSchema),
            want: "50 %"
        },
        {
            desc: "Progress Bar without score and scorelimit",
            submission: create(SubmissionSchema),
            assignment: create(AssignmentSchema),
            want: "0 %"
        },
        {
            desc: "Progress Bar with incorrect index",
            submission: create(SubmissionSchema, { score: 50 }),
            assignment: create(AssignmentSchema, { scoreLimit: 100 }),
            want: "0 %",
            assignmentIndex: 10
        }
    ]

    test.each(progressBarTests)(`[Progress.LAB] $desc`, (test) => {
        labTest(test, true)
    })

    test.each(progressBarTests)(`[Progress.LAB - Without Submission] $desc`, (test) => {
        labTest(test, false)
    })


    test.each(progressBarTests)(`[Progress.NAV] $desc`, (test) => {
        const overmind = initializeOvermind({ assignments: { "1": [test.assignment] }, submissions: { "1": [test.submission] } })

        const { container } = render(
            <Provider value={overmind}>
                <ProgressBar courseID={"1"} assignmentIndex={0} submission={test.submission} type={Progress.NAV} />
            </Provider>
        )


        const bar = container.getElementsByTagName("div").item(0)
        expect(bar?.style).toHaveProperty("right", `${100 - test.submission.score}%`)
        expect(bar?.style).toHaveProperty(
            "border-bottom",
            test.submission.score >= test.assignment.scoreLimit
                ? "2px solid green"
                : "2px solid yellow"
        )
    })
})

const labTest = (test: ProgressBarTest, withSubmission: boolean) => {
    const overmind = initializeOvermind({ assignments: { "1": test.assignment ? [test.assignment] : [] }, submissions: { "1": test.submission ? [test.submission] : [] } })

    const { container } = render(
        <Provider value={overmind}>
            <ProgressBar courseID={"1"} assignmentIndex={test.assignmentIndex ?? 0} submission={withSubmission ? test.submission : undefined} type={Progress.LAB} />
        </Provider>
    )

    // Incorrect assignment index should not have a secondary bar
    const hasSecondary = test.submission.score < test.assignment.scoreLimit && test.assignmentIndex === undefined
    // Given an invalid assignment index, we expect the bar to be empty
    // However, if we pass a submission, we expect the bar to be filled to the correct percentage
    const score = test.assignmentIndex === undefined || withSubmission
        ? test.submission.score
        : 0

    const bars = container.getElementsByClassName("progress-bar")
    expect(bars).toHaveLength(hasSecondary ? 2 : 1)
    if (hasSecondary) {
        const secondary = container.getElementsByClassName("progressbar-secondary").item(0)
        if (!secondary) {
            fail()
        }
        expect(secondary.getAttribute("style")).toContain(`width: ${test.assignment.scoreLimit - test.submission.score}%`)
        expect(secondary.textContent).toEqual(`${test.assignment.scoreLimit - test.submission.score} %`)
    }

    expect(container.getElementsByClassName("progress-bar bg-primary").item(0)?.getAttribute("style")).toContain(`width: ${score}%`)
    expect(bars[0].textContent).toContain(test.want)
}
