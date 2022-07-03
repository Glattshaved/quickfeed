import React, { Dispatch, SetStateAction, useMemo } from "react"
import { hasEnrollment } from "../../Helpers"
import { useActions, useAppState } from "../../overmind"
import FormInput from "../forms/FormInput"
import { useHistory } from "react-router"
import { Converter } from "../../convert"


const ProfileForm = ({ children, setEditing }: { children: React.ReactNode, setEditing: Dispatch<SetStateAction<boolean>> }): JSX.Element => {
    const state = useAppState()
    const actions = useActions()
    const history = useHistory()

    const signup = useMemo(() => !state.isValid, [state.isValid])

    // Create a copy of the user object, so that we can modify it without affecting the original object.
    const user = Converter.clone(state.self)

    // Update the user object when user input changes, and update the state.
    const handleChange = (event: React.FormEvent<HTMLInputElement>) => {
        const { name, value } = event.currentTarget
        switch (name) {
            case "name":
                user.name = value
                break
            case "email":
                user.email = value
                break
            case "studentid":
                user.studentid = value
                break
        }
    }


    // Sends the updated user object to the server on submit.
    const submitHandler = () => {
        actions.updateUser(user)
        // Disable editing after submission
        setEditing(false)
        if (!hasEnrollment(state.enrollments)) {
            history.push("/courses")
        }
    }

    return (
        <div>
            {signup ? children : null}
            <form className="form-group" onSubmit={e => { e.preventDefault(); submitHandler() }}>
                <FormInput prepend="Name" name="name" defaultValue={user.name} onChange={handleChange} />
                <FormInput prepend="Email" name="email" defaultValue={user.email} onChange={handleChange} type="email" />
                <FormInput prepend="Student ID" name="studentid" defaultValue={user.studentid} onChange={handleChange} type="number" />
                <div className="col input-group mb-3">
                    <input className="btn btn-primary" disabled={!state.isValid} type="submit" value="Save" style={{ marginTop: "20px" }} />
                </div>
            </form>
        </div>
    )
}

export default ProfileForm
