import React, { useState } from "react"
import { Link, useHistory } from "react-router-dom"
import { useActions, useAppState } from "../../overmind"


const NavBarFooter = () => {
    const logout = useActions().logout
    const self = useAppState((state) => state.self)

    const history = useHistory()

    const [hidden, setHidden] = useState<boolean>(true)

    const LoginButton = () => {
        if (self.getId() > 0) {
            return (
            <li hidden={hidden}>
                <a href="/logout" className="Sidebar-items-link" onClick={() => logout()}>Log out</a>
            </li>
            )
        }
        return (
            <li>
                <a href="/auth/github" style={{textAlign:"center",paddingTop:"15px"}}>
                    <i className="fa fa-2x fa-github" id="github"></i>
                </a>
            </li>
        )
    }

    const ProfileButton = () => {
        if (self.getId() > 0) {
            return (
                <li key="profile" onClick={() => history.push("/profile")} onMouseEnter={() => setHidden(false) }>
                    <div><img src={self.getAvatarurl()} id="avatar"></img></div>    
                </li>
            )
        }
        return null
    }

    const AboutButton = () => {
        return (
            <li key="about" hidden={hidden}>
                <Link to="/about" className="Sidebar-items-link">
                    About
                </Link>
            </li>
        )
    }

    const AdminButton = () => {
        if (self.getIsadmin()) {
            return (           
            <li key="admin" hidden={hidden}>
                <Link to="/admin" className="Sidebar-items-link">
                    Admin
                </Link>
            </li>
            )
        }
        return null
        
    }

    return (
        <div className="SidebarFooter" onMouseLeave={() => setHidden(true)}>            
            <AboutButton />
            <AdminButton />
            <LoginButton />
            <ProfileButton />
        </div>
    )
}

export default NavBarFooter