import React, { useEffect } from "react"
import { useHistory } from "react-router"


const RedirectButton = ({to}: {to: string}) => {
    const history = useHistory()
    const hide = history.location.pathname == to ? true : false
    
    useEffect(() => {
        const handleKeyPress = (e: KeyboardEvent) => {
            if (e.key == "Backspace" && !hide) {
                history.push(to)
            }
        }

        document.addEventListener('keydown', handleKeyPress)

        const cleanup = () => {
            document.removeEventListener('keydown', handleKeyPress);
        }
        return cleanup
    }, [])

    return (
        <div className={"btn btn-dark redirectButton"} onClick={() => history.push(to)} hidden={hide}>
            <i className="fa fa-arrow-left"></i>
        </div>
    )
}

export default RedirectButton