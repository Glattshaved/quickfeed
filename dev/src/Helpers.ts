import { Assignment, Course, Enrollment, EnrollmentLink, GradingBenchmark, GradingCriterion, Group, Review, Submission, SubmissionLink, User } from "../proto/ag/ag_pb"
import { BuildInfo, Score } from "../proto/kit/score/score_pb"
import { useParams } from "react-router"


export enum Color {
    RED = "danger",
    BLUE = "primary",
    GREEN = "success",
    YELLOW = "warning",
    GRAY = "secondary",
    WHITE = "light",
    BLACK = "dark",
}

export enum Sort {
    NAME,
    STATUS,
    ID
}

/** Returns a string with a prettier format for a deadline */
export const getFormattedTime = (deadline_string: string): string => {
    const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
    const deadline = new Date(deadline_string)
    return `${deadline.getDate()} ${months[deadline.getMonth()]} ${deadline.getFullYear()} ${deadline.getHours()}:${deadline.getMinutes() < 10 ? "0" + deadline.getMinutes() : deadline.getMinutes()}`
}

export interface Deadline {
    className: string,
    message: string,
    daysUntil: number,
}

/** Utility function for LandingpageTable functionality. To format the output string and class/css based on how far the deadline is in the future */
// layoutTime = "2021-03-20T23:59:00"
export const timeFormatter = (deadline: string): Deadline => {
    const timeToDeadline = new Date(deadline).getTime() - new Date().getTime()
    const days = Math.floor(timeToDeadline / (1000 * 3600 * 24))
    const hours = Math.floor(timeToDeadline / (1000 * 3600))
    const minutes = Math.floor((timeToDeadline % (1000 * 3600)) / (1000 * 60))

    if (timeToDeadline < 0) {
        return { className: "table-danger", message: `Expired ${-days > 0 ? -days + " days ago" : -hours + " hours"}`, daysUntil: 0 }
    }

    if (days == 0) {
        return { className: "table-danger", message: `${hours} hours and ${minutes} minutes to deadline!`, daysUntil: 0 }
    }

    if (days < 3) {
        return { className: "table-warning", message: `${days} day${days == 1 ? " " : "s"} to deadline`, daysUntil: days }
    }

    if (days < 14) {
        return { className: "table-primary", message: `${days} days`, daysUntil: days }
    }

    return { className: "", message: "", daysUntil: days }
}

// Used for displaying enrollment status
export const EnrollmentStatus = {
    0: "None",
    1: "Pending",
    2: "Student",
    3: "Teacher",
}

/*
    arr: Any array, ex. Enrollment[], User[],
    funcs: an array of functions that will be applied in order to reach the field to sort on
    by: A function returning an element to sort on

    Example:
        To sort state.enrollmentsByCourseId[2].getUser().getName() by name, call like
        (state.enrollmentsByCourseId[2], [Enrollment.prototype.getUser], User.prototype.getName)

    Returns an array of the same type as arr, sorted by the by-function
*/
export const sortByField = (arr: any[], funcs: Function[], by: Function, descending?: boolean) => {
    const unsortedArray = Object.assign([], arr)
    const sortedArray = unsortedArray.sort((a, b) => {
        let x: any
        let y: any
        if (!a || !b) {
            return 0
        }
        if (funcs.length > 0) {
            funcs.forEach(func => {
                x = x ? func.call(x) : func.call(a)
                y = y ? func.call(y) : func.call(b)
            })
        } else {
            x = a
            y = b
        }
        if (by.call(x) === by.call(y)) {
            return 0
        }
        if (by.call(x) < by.call(y)) {
            return descending ? 1 : -1
        }
        if (by.call(x) > by.call(y)) {
            return descending ? -1 : 1
        }
        return -1
    })
    return sortedArray
}

// TODO: Could be computed on the backend (https://github.com/quickfeed/quickfeed/issues/420)
/** getPassedTestCount returns a string with the number of passed tests and the total number of tests */
export const getPassedTestsCount = (score: Score.AsObject[]): string => {
    let totalTests = 0
    let passedTests = 0
    score.forEach(score => {
        if (score.score === score.maxscore) {
            passedTests++
        }
        totalTests++
    })
    if (totalTests === 0) {
        return ""
    }
    return `${passedTests}/${totalTests}`
}

export const isValid = (elm: User | EnrollmentLink): boolean => {
    if (elm instanceof User) {
        return elm.getName().length > 0 && elm.getEmail().length > 0 && elm.getStudentid().length > 0
    }
    if (elm instanceof EnrollmentLink) {
        return elm.getEnrollment()?.getUser() !== undefined && elm.getSubmissionsList().length > 0
    }
    return true
}

/** hasEnrollment returns true if any of the provided has been approved */
export const hasEnrollment = (enrollments: Enrollment.AsObject[]): boolean => {
    return enrollments.some(enrollment => enrollment.status > Enrollment.UserStatus.PENDING)
}

export const isStudent = (enrollment: Enrollment.AsObject): boolean => { return hasStudent(enrollment.status) }
export const isTeacher = (enrollment: Enrollment.AsObject): boolean => { return hasTeacher(enrollment.status) }
export const isPending = (enrollment: Enrollment.AsObject): boolean => { return hasPending(enrollment.status) }

export const isPendingGroup = (group: Group.AsObject): boolean => { return group.status === Group.GroupStatus.PENDING }
export const isApprovedGroup = (group: Group.AsObject): boolean => { return group.status === Group.GroupStatus.APPROVED }

/** isEnrolled returns true if the user is enrolled in the course, and is no longer pending. */
export const isEnrolled = (enrollment: Enrollment.AsObject): boolean => { return enrollment.status >= Enrollment.UserStatus.STUDENT }

/** toggleUserStatus switches between teacher and student status. */
export const toggleUserStatus = (enrollment: Enrollment.AsObject): Enrollment.UserStatus => {
    return isTeacher(enrollment) ? Enrollment.UserStatus.STUDENT : Enrollment.UserStatus.TEACHER
}

export const hasNone = (status: Enrollment.UserStatus): boolean => { return status === Enrollment.UserStatus.NONE }
export const hasPending = (status: Enrollment.UserStatus): boolean => { return status === Enrollment.UserStatus.PENDING }
export const hasStudent = (status: Enrollment.UserStatus): boolean => { return status === Enrollment.UserStatus.STUDENT }
export const hasTeacher = (status: Enrollment.UserStatus): boolean => { return status === Enrollment.UserStatus.TEACHER }

/** hasEnrolled returns true if user has enrolled in course, or is pending approval. */
export const hasEnrolled = (status: Enrollment.UserStatus): boolean => { return status >= Enrollment.UserStatus.PENDING }

export const isVisible = (enrollment: Enrollment.AsObject): boolean => { return enrollment.state === Enrollment.DisplayState.VISIBLE }
export const isFavorite = (enrollment: Enrollment.AsObject): boolean => { return enrollment.state === Enrollment.DisplayState.FAVORITE }

export const isCourseCreator = (user: User.AsObject, course: Course.AsObject): boolean => { return user.id === course.coursecreatorid }
export const isAuthor = (user: User.AsObject, review: Review.AsObject): boolean => { return user.id === review.reviewerid }

export const isManuallyGraded = (assignment: Assignment.AsObject): boolean => {
    return assignment.reviewers > 0
}

export const isApproved = (submission: Submission.AsObject): boolean => { return submission.status === Submission.Status.APPROVED }
export const isRevision = (submission: Submission.AsObject): boolean => { return submission.status === Submission.Status.REVISION }
export const isRejected = (submission: Submission.AsObject): boolean => { return submission.status === Submission.Status.REJECTED }

export const hasReviews = (submission: Submission.AsObject): boolean => { return submission.reviewsList.length > 0 }
export const hasBenchmarks = (obj: Review.AsObject | Assignment.AsObject): boolean => { return obj.gradingbenchmarksList.length > 0 }
export const hasCriteria = (benchmark: GradingBenchmark.AsObject): boolean => { return benchmark.criteriaList.length > 0 }
export const hasEnrollments = (obj: Group.AsObject): boolean => { return obj.enrollmentsList.length > 0 }

/** getCourseID returns the course ID determined by the current route */
export const getCourseID = (): number => {
    const route = useParams<{ id?: string }>()
    return Number(route.id)
}

export const isHidden = (value: string, query: string): boolean => {
    return !value.toLowerCase().includes(query) && query.length > 0
}

/** getSubmissionsScore calculates the total score of all submissions in a SubmissionLink[] */
export const getSubmissionsScore = (submissions: SubmissionLink.AsObject[]): number => {
    let score = 0
    submissions.forEach(link => {
        if (!link.submission) {
            return
        }
        score += link.submission.score
    })
    return score
}

/** getNumApproved returns the number of approved submissions in a SubmissionLink[] */
export const getNumApproved = (submissions: SubmissionLink.AsObject[]): number => {
    let num = 0
    submissions.forEach(submission => {
        if (!submission.submission) {
            return
        }
        if (isApproved(submission.submission)) {
            num++
        }
    })
    return num
}

export const getSubmissionByAssignmentID = (submissions: SubmissionLink.AsObject[] | undefined, assignmentID: number): Submission.AsObject | undefined => {
    return submissions?.find(submission => submission.assignment?.id === assignmentID)?.submission
}

export const EnrollmentStatusBadge = {
    0: "",
    1: "badge badge-info",
    2: "badge badge-primary",
    3: "badge badge-danger",
}

/** SubmissionStatus returns a string with the status of the submission, given the status number, ex. Submission.Status.APPROVED -> "Approved" */
export const SubmissionStatus = {
    0: "None",
    1: "Approved",
    2: "Rejected",
    3: "Revision",
}

// TODO: This could possibly be done on the server. Would need to add a field to the proto submission/score model.
/** assignmentStatusText returns a string that is used to tell the user what the status of their submission is */
export const assignmentStatusText = (assignment: Assignment.AsObject, submission: Submission.AsObject): string => {
    // If the submission is not graded, return a descriptive text
    if (submission.status === Submission.Status.NONE) {
        // If the assignment requires manual approval, and the score is above the threshold, return Await Approval
        if (!assignment.autoapprove && submission.score >= assignment.scorelimit) {
            return "Awaiting approval"
        }
        if (submission.score < assignment.scorelimit) {
            return `Need ${assignment.scorelimit}% score for approval`
        }
    }
    // If the submission is graded, return the status
    return SubmissionStatus[submission.status]
}

// Helper functions for default values for new courses
export const defaultTag = (date: Date): string => {
    return date.getMonth() >= 10 || date.getMonth() < 4 ? "Spring" : "Fall"
}

export const defaultYear = (date: Date): number => {
    return (date.getMonth() <= 11 && date.getDate() <= 31) && date.getMonth() > 10 ? (date.getFullYear() + 1) : date.getFullYear()
}

export const userLink = (user: User.AsObject): string => {
    return `https://github.com/${user.login}`
}

export const userRepoLink = (course: Course.AsObject, user: User.AsObject): string => {
    return `https://github.com/${course.organizationpath}/${user.login}-labs`
}

export const groupRepoLink = (course: Course.AsObject, group: Group.AsObject): string => {
    course.organizationpath
    return `https://github.com/${course.organizationpath}/${slugify(group.name)}`
}

export const getSubmissionCellColor = (submission: Submission.AsObject): string => {
    if (isApproved(submission)) {
        return "result-approved"
    }
    if (isRevision(submission)) {
        return "result-revision"
    }
    if (isRejected(submission)) {
        return "result-rejected"
    }
    return "clickable"
}

const slugify = (str: string): string => {
    str = str.replace(/^\s+|\s+$/g, "").toLowerCase()

    // Remove accents, swap ñ for n, etc
    const from = "ÁÄÂÀÃÅČÇĆĎÉĚËÈÊẼĔȆÍÌÎÏŇÑÓÖÒÔÕØŘŔŠŤÚŮÜÙÛÝŸŽáäâàãåčçćďéěëèêẽĕȇíìîïňñóöòôõøðřŕšťúůüùûýÿžþÞĐđßÆaæ·/,:;&"
    const to = "AAAAAACCCDEEEEEEEEIIIINNOOOOOORRSTUUUUUYYZaaaaa-cccdeeeeeeeeiiiinnooooo-orrstuuuuuyyzbBDdBAa-------"
    for (let i = 0; i < from.length; i++) {
        str = str.replace(new RegExp(from.charAt(i), "g"), to.charAt(i))
    }

    // Remove invalid chars, replace whitespace by dashes, collapse dashes
    return str.replace(/[^a-z0-9 -_]/g, "").replace(/\s+/g, "-").replace(/-+/g, "-")
}

/* Use this function to simulate a delay in the loading of data */
/* Used in development to simulate a slow network connection */
export const delay = (ms: number) => {
    return new Promise(resolve => setTimeout(resolve, ms))
}


export enum EnrollmentSort {
    Name,
    Status,
    Email,
    Activity,
    Slipdays,
    Approved,
    StudentID
}

export enum SubmissionSort {
    Name,
    Status,
    Score,
    Approved
}

/** Sorting */
const enrollmentCompare = (a: Enrollment.AsObject, b: Enrollment.AsObject, sortBy: EnrollmentSort, descending: boolean): number => {
    const m = descending ? -1 : 1
    switch (sortBy) {
        case EnrollmentSort.Name:
            const nameA = a.user?.name ?? ""
            const nameB = b.user?.name ?? ""
            return m * (nameA.localeCompare(nameB))
        case EnrollmentSort.Status:
            return m * (a.status - b.status)
        case EnrollmentSort.Email:
            const emailA = a.user?.email ?? ""
            const emailB = b.user?.email ?? ""
            return m * (emailA.localeCompare(emailB))
        case EnrollmentSort.Activity:
            return m * (new Date(a.lastactivitydate).getTime() - new Date(b.lastactivitydate).getTime())
        case EnrollmentSort.Slipdays:
            return m * (a.slipdaysremaining - b.slipdaysremaining)
        case EnrollmentSort.Approved:
            return m * (a.totalapproved - b.totalapproved)
        case EnrollmentSort.StudentID:
            const aID = a.user?.id ?? 0
            const bID = b.user?.id ?? 0
            return m * (aID - bID)
        default:
            return 0
    }
}

export const sortEnrollments = (enrollments: Enrollment.AsObject[], sortBy: EnrollmentSort, descending: boolean): Enrollment.AsObject[] => {
    return enrollments.sort((a, b) => {
        return enrollmentCompare(a, b, sortBy, descending)
    })
}

// Class with converter functions for the different proto types
// TODO(jostein): Move this to a separate file
// TODO(jostein): Currently used in `actions`. Consider moving use to `GRPCManager`.
export class ProtoConverter {

    /**
     * create creates a new object of the given type
     * @param type the type you wish to create as T.AsObject
     * @returns the given type as a T.AsObject
     * @example create<User.AsObject>(User) // returns a new User.AsObject
     */
    public static create<T>(type: (new () => any)): T {
        return (new type()).toObject() as T
    }

    public static clone<T>(obj: T): T {
        return JSON.parse(JSON.stringify(obj)) as T
    }

    /** The following functions will convert various kinds of objects
     *  into their respective protobuf message types */
    // TODO(jostein): Add converting functions for all types
    // TODO(jostein): It would be awesome to make this generic. Not sure how to do that though.
    public static toUser = (obj: User.AsObject): User => {
        const user = new User()
        user.setId(obj.id)
        user.setName(obj.name)
        user.setEmail(obj.email)
        user.setIsadmin(obj.isadmin)
        user.setLogin(obj.login)
        user.setAvatarurl(obj.avatarurl)
        user.setStudentid(obj.studentid)
        const enrollments = obj.enrollmentsList.map(e => this.toEnrollment(e))
        user.setEnrollmentsList(enrollments)

        return user
    }

    public static toEnrollment = (obj: Enrollment.AsObject): Enrollment => {
        const enrollment = new Enrollment()
        enrollment.setId(obj.id)
        enrollment.setStatus(obj.status)
        enrollment.setLastactivitydate(obj.lastactivitydate)
        enrollment.setSlipdaysremaining(obj.slipdaysremaining)
        enrollment.setTotalapproved(obj.totalapproved)
        enrollment.setCourseid(obj.courseid)
        enrollment.setUserid(obj.userid)
        enrollment.setGroupid(obj.groupid)
        enrollment.setHasteacherscopes(obj.hasteacherscopes)
        enrollment.setState(obj.state)
        for (const slipdays of obj.usedslipdaysList) {
            // Handle usedslipdays
        }
        if (obj.user) {
            enrollment.setUser(this.toUser(obj.user))
        }
        if (obj.course) {
            enrollment.setCourse(this.toCourse(obj.course))
        }
        if (obj.group) {
            enrollment.setGroup(this.toGroup(obj.group))
        }
        return enrollment
    }

    public static toCourse = (obj: Course.AsObject): Course => {
        const course = new Course()
        course.setId(obj.id)
        course.setCoursecreatorid(obj.coursecreatorid)
        course.setName(obj.name)
        course.setCode(obj.code)
        course.setYear(obj.year)
        course.setTag(obj.tag)
        course.setProvider(obj.provider)
        course.setOrganizationid(obj.organizationid)
        course.setOrganizationpath(obj.organizationpath)
        course.setSlipdays(obj.slipdays)
        course.setDockerfile(obj.dockerfile)
        course.setEnrolled(obj.enrolled)

        const enrollments = obj.enrollmentsList.map(e => this.toEnrollment(e))
        course.setEnrollmentsList(enrollments)

        const assignments = obj.assignmentsList.map(a => this.toAssignment(a))
        course.setAssignmentsList(assignments)

        const groups = obj.groupsList.map(g => this.toGroup(g))
        course.setGroupsList(groups)

        return course
    }

    public static toAssignment = (obj: Assignment.AsObject): Assignment => {
        const assignment = new Assignment()
        assignment.setId(obj.id)
        assignment.setCourseid(obj.courseid)
        assignment.setName(obj.name)
        assignment.setScriptfile(obj.scriptfile)
        assignment.setDeadline(obj.deadline)
        assignment.setAutoapprove(obj.autoapprove)
        assignment.setOrder(obj.order)
        assignment.setIsgrouplab(obj.isgrouplab)
        assignment.setScorelimit(obj.scorelimit)
        assignment.setReviewers(obj.reviewers)

        const submissions = obj.submissionsList.map(s => this.toSubmission(s))
        assignment.setSubmissionsList(submissions)

        const gradingBenchmarks = obj.gradingbenchmarksList.map(g => this.toGradingBenchmark(g))
        assignment.setGradingbenchmarksList(gradingBenchmarks)

        assignment.setContainertimeout(obj.containertimeout)

        return assignment
    }

    public static toSubmission = (obj: Submission.AsObject): Submission => {
        const submission = new Submission()
        submission.setId(obj.id)
        submission.setAssignmentid(obj.assignmentid)
        submission.setUserid(obj.userid)
        submission.setGroupid(obj.groupid)
        submission.setScore(obj.score)
        submission.setCommithash(obj.commithash)
        submission.setReleased(obj.released)
        submission.setStatus(obj.status)
        submission.setApproveddate(obj.approveddate)

        const reviews = obj.reviewsList.map(r => this.toReview(r))
        submission.setReviewsList(reviews)

        if (obj.buildinfo) {
            submission.setBuildinfo(this.toBuildInfo(obj.buildinfo))
        }

        const scores = obj.scoresList.map(s => this.toScore(s))
        submission.setScoresList(scores)

        return submission
    }

    public static toBuildInfo = (obj: BuildInfo.AsObject): BuildInfo => {
        const buildInfo = new BuildInfo()
        buildInfo.setId(obj.id)
        buildInfo.setSubmissionid(obj.submissionid)
        buildInfo.setBuilddate(obj.builddate)
        buildInfo.setBuildlog(obj.buildlog)
        buildInfo.setExectime(obj.exectime)

        return buildInfo
    }

    public static toScore = (obj: Score.AsObject): Score => {
        const score = new Score()
        score.setId(obj.id)
        score.setSubmissionid(obj.submissionid)
        score.setSecret(obj.secret)
        score.setTestname(obj.testname)
        score.setScore(obj.score)
        score.setMaxscore(obj.maxscore)
        score.setWeight(obj.weight)
        score.setTestdetails(obj.testdetails)

        return score
    }

    public static toReview = (obj: Review.AsObject): Review => {
        const review = new Review()
        review.setId(obj.id)
        review.setSubmissionid(obj.submissionid)
        review.setReviewerid(obj.reviewerid)
        review.setFeedback(obj.feedback)
        review.setReady(obj.ready)
        review.setScore(obj.score)

        const gradingBenchmarks = obj.gradingbenchmarksList.map(g => this.toGradingBenchmark(g))
        review.setGradingbenchmarksList(gradingBenchmarks)

        review.setEdited(obj.edited)

        return review
    }

    public static toGradingBenchmark = (obj: GradingBenchmark.AsObject): GradingBenchmark => {
        const gradingBenchmark = new GradingBenchmark()
        gradingBenchmark.setId(obj.id)
        gradingBenchmark.setAssignmentid(obj.assignmentid)
        gradingBenchmark.setReviewid(obj.reviewid)
        gradingBenchmark.setHeading(obj.heading)
        gradingBenchmark.setComment(obj.comment)

        const criteria = obj.criteriaList.map(c => this.toGradingCriterion(c))
        gradingBenchmark.setCriteriaList(criteria)

        return gradingBenchmark
    }

    public static toGradingCriterion = (obj: GradingCriterion.AsObject): GradingCriterion => {
        const criterion = new GradingCriterion()
        criterion.setId(obj.id)
        criterion.setBenchmarkid(obj.benchmarkid)
        criterion.setPoints(obj.points)
        criterion.setDescription(obj.description)
        criterion.setGrade(obj.grade)
        criterion.setComment(obj.comment)

        return criterion
    }

    public static toGroup = (obj: Group.AsObject): Group => {
        const group = new Group()
        group.setId(obj.id)
        group.setName(obj.name)
        group.setCourseid(obj.courseid)
        group.setTeamid(obj.teamid)
        group.setStatus(obj.status)

        const users = obj.usersList.map(u => this.toUser(u))
        group.setUsersList(users)

        const enrollments = obj.enrollmentsList.map(e => this.toEnrollment(e))
        group.setEnrollmentsList(enrollments)

        return group
    }

    public static load = () => {
        const data = localStorage.getItem("state")
        if (data) {
            const obj = JSON.parse(data)
            console.log("Loaded state: ", obj)
            return obj
        }
        return undefined
    }

}
