import React, { useEffect } from "react"

const LoginPage = (): JSX.Element => {
  // Add a class to the body element to style the login page
  useEffect(() => {
    document.body.classList.add("login-page")
    return () => {
      document.body.classList.remove("login-page")
    }
  }, [])
  return (
    <div className="loginContainer">
      <h1 className="loginWelcomeHeader">Welcome to QuickFeed</h1>
      <p className="lead mt-5 mb-5" style={{ textAlign: "center", marginBottom: "50px" }}>
        To get started with QuickFeed, please sign in with your GitHub account.
      </p>
      <section id="loginBox">
        <div className="mb-5 loginBox">
          <i className="fa fa-5x fa-github align-middle ms-auto mb-3" id="github icon"></i>
          <h4>Sign in with GitHub</h4>
          <p className="text-secondary"> to continue to QuickFeed </p>
          <a href="/auth/github" className="loginButton"> Sign in </a>
        </div>
      </section>
      <div className="container">
        <section id="aboutquickfeed">
          {/* <hr className="loginDivider" /> */}
          <h2 className="featurette-heading">About QuickFeed</h2>
          <p className="lead">
            QuickFeed is a tool for providing automated feedback to students on their lab assignments.
            QuickFeed builds upon version control systems and continuous integration.
            When students upload code to their repositories, QuickFeed automatically builds their code and provides feedback based on tests supplied by the teaching staff.
            When grading assignments, teaching staff can access the results of test execution and have a valuable tool in the grading process.
          </p>
          <hr className="loginDivider" />
        </section>
        <div key="rowheader" className="row marketing">
          <div key="gh" className="col-lg-4">
            <i
              className="fa fa-github align-middle ms-auto"
              id="github logo"
              style={{ width: "140px", height: "140px", fontSize: "140px", lineHeight: "140px" }} />
            <h2>GitHub Integration</h2>
            <p>
              Manage all students and courses on GitHub.
              Each student gets their own repository.
              Teachers get separate repositories for publishing assignments and information to students.
              All taken care of automatically.
            </p>
          </div>
          <div key="ci" className="col-lg-4">
            <img
              className="img-circle"
              src="/assets/img/overlapping-arrows-no-background.png"
              style={{ width: "140px", height: "140px" }}
            />
            <h2>Continuous Integration</h2>
            <p>
              Instantaneous feedback to students on how well their code performs.
              Students can quickly identify what they need to focus on to improve.
              All customizable for the teaching staff.
            </p>
          </div>
          <div key="grade" className="col-lg-4">
            <img
              className="img-circle"
              src="/assets/img/Aplus2-no-background.png"
              alt="A+ image" style={{ width: "140px", height: "140px" }} />
            <h2>Fair Grading</h2>
            <p>
              On due date of an assignment, the most recent version
              of each student&apos;s code is available through GitHub.
              Easily accessible for the teachers.
              Together with latest build log, this makes grading easier and more fair.
            </p>
          </div>
        </div>
        <section id="automatedfeedback">
          <hr className="loginDivider" />
          <div key="row1" className="row featurette">
            <div key="c1r1" className="col-md-7">
              <h2 className="featurette-heading">
                QuickFeed: <span className="text-muted">Automated student feedback</span>
              </h2>
              <p className="lead">
                QuickFeed aims to provide students with fast
                feedback on their lab assignments, and is designed
                to help students learn about state-of-the-art tools
                used in the industry.
                QuickFeed builds upon version control systems and
                continuous integration.
                When students upload code to their repositories,
                QuickFeed automatically builds their code and
                provides feedback based on tests supplied by the
                teaching staff.
                When grading assignments, teaching staff can access
                the results of test execution and have a valuable
                tool in the grading process.
              </p>
            </div>
            <div key="c2r1" className="col-md-5">
              <img
                className="featurette-image img-responsive about"
                src="/assets/img/intro1.png"
                alt="Generic placeholder image" />
            </div>
          </div>
        </section>
        <section id="versioncontrol">
          <hr className="loginDivider" />
          <div key="row2" className="row featurette">
            <div key="c1r2" className="col-md-5">
              <img
                className="featurette-image img-responsive about"
                src="/assets/img/intro3.png"
                alt="Generic placeholder image" />
            </div>
            <div key="c2r2" className="col-md-7">
              <h2 className="featurette-heading">
                GitHub Integration: <span className="text-muted">Managing courses and students</span>
              </h2>
              <p className="lead">
                A course is an organization on GitHub.
                Students get access to their own private GitHub repository.
                Uploading their code for review or grading, students can
                learn to use git for version control.
              </p>
            </div>
          </div>
        </section>
        <section id="ci">
          <hr className="loginDivider" />
          <div key="row3" className="row featurette">
            <div key="c1r3" className="col-md-7">
              <h2 className="featurette-heading">
                Continuous Integration: <span className="text-muted">builds and tests student code. </span>
              </h2>
              <p className="lead">
                As code gets pushed up to GitHub, an automatic build process
                defined by the teacher, generates feedback to students.
                When the build process is completed, student gets immediate
                access to this feedback on their personal course page.
                Tests defined by either teachers or students will be processed
                and tell students about their progress on the assignments.
              </p>
            </div>
            <div key="c2r3" className="col-md-5">
              <img
                className="featurette-image img-responsive about"
                src="/assets/img/intro2.png"
                alt="Generic placeholder image" />
            </div>
          </div>
        </section>
        <section id="grading">
          <hr className="loginDivider" />
          <div key="row4" className="row featurette">
            <div key="c1r4" className="col-md-5">
              <img
                className="featurette-image img-responsive about"
                src="/assets/img/intro4.png"
                alt="Generic placeholder image" />
            </div>
            <div key="c2r4" className="col-md-7">
              <h2 className="featurette-heading">
                Grading: <span className="text-muted">Easy and Fair</span>
              </h2>
              <p className="lead">
                On the due date, teachers can access the test results and
                use this as a tool in the grading process.
                The teaching staff will immediately know which of their
                tests passed, and how much of the code is covered by the tests.
              </p>
            </div>
          </div>
        </section>
        <footer className="text-center mt-5">
          <button
            onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}
            className="btn align-items-center backToTop"
          >
            <i className="fa fa-arrow-up" />
            <p>Back to top</p>
          </button>
        </footer>
      </div>
    </div>
  )
}

export default LoginPage
