import React from "react"
import { Enrollment, User as pbUser } from "../../../proto/qf/types_pb"
import { useGrpc } from "../../overmind"
import { EnrollmentStatus, EnrollmentStatusBadge } from "../../Helpers"

const User = ({ user }: { user: pbUser; hidden: boolean }): JSX.Element => {
    const { api } = useGrpc()
    const [enrollments, setEnrollments] = React.useState<Enrollment[]>([])
    const [showEnrollments, setShowEnrollments] = React.useState<boolean>(false)

    const toggleEnrollments = () => {
        setShowEnrollments(!showEnrollments)
        if (!enrollments.length) {
            console.log("Getting enrollments", user)
            getEnrollments()
        }
    }

    const getEnrollments = () => {
        api.client
            .getEnrollments({
                FetchMode: { case: "userID", value: user.ID },
            })
            .then((response) => {
                setEnrollments(response.message.enrollments)
            })
    }

    return (
        <div className="clickable" onClick={toggleEnrollments}>
            {user.Name}
            {user.IsAdmin ? (
                <span className={"badge badge-primary ml-2"}>Admin</span>
            ) : null}

            {enrollments.length && showEnrollments ? (
                <div className="ml-4">
                    {enrollments.map((enrollment) => (
                        <div key={enrollment.ID.toString()}>
                            <span className="badge badge-secondary">
                                {enrollment.course?.name}
                            </span>{" "}
                            <span
                                className={
                                    EnrollmentStatusBadge[enrollment.status]
                                }
                            >
                                {EnrollmentStatus[enrollment.status]}
                            </span>
                        </div>
                    ))}
                </div>
            ) : null}
        </div>
    )
}

export default User
