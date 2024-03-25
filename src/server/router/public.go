package router

import (
	"github.com/zayarhtet/seap-api/src/server/controller"
)

func publicRoutes() {
	public := seapRouter.Group("/")
	public.GET("/", controller.Welcome())

	authed := seapRouter.Group("/api/auth")
	authed.POST("/register", controller.Register())
	authed.POST("/login", controller.Login())
}

func adminRoutes() {
	admin := seapRouter.Group("/api/admin/")
	admin.Use(controller.AdminMiddleware())
	admin.GET("/roles", controller.GetAllRoles())
	admin.GET("/role/:id", controller.GetRoleById())
	admin.GET("/members", controller.GetAllMembers())
	admin.GET("/members/families", controller.GetAllMembersWithFamilies())
	admin.GET("/member/:id", controller.GetMemberById())
	admin.GET("/member/:id/families", controller.GetMemberByIdWithFamilies())
	admin.DELETE("/member/:id", controller.DeleteMember())
	admin.GET("/families", controller.GetAllFamilies())
	admin.GET("/families/members", controller.GetAllFamiliesWithMembers())
	admin.GET("/duties", controller.GetAllDuties())
	//admin.GET("/grading", controller.GetAllGrading())
}

func individualRoutes() {
	protected := seapRouter.Group("/api/my/")
	protected.Use(controller.IndividualMiddleware())
	protected.GET("/role", controller.GetMyRole())
	protected.GET("/member", controller.GetMyMember())
	protected.GET("/families", controller.GetMyFamilies())
	protected.GET("/duties", controller.GetMyDuties()) // get all duties associated with the user
}

func tutorRoutes() {
	tutor := seapRouter.Group("/api/my/")
	tutor.Use(controller.TutorMiddleware())
	tutor.POST("/create/family", controller.SaveNewFamily())
}

func familyTutorRoutes() {
	// execute only if the user has a tutor role in the family
	familyTutor := seapRouter.Group("/api/my/")
	familyTutor.Use(controller.FamilyTutorMiddleware())
	familyTutor.POST("/family/:famId/addMember", controller.AddNewMemberToFamily())
	familyTutor.GET("/family/:famId/members", controller.GetAllMembersByFamilyId()) // get all duties associated with the fam id
	//familyTutor.POST("family/:famId/create/grade", controller.AddNewGrade())        // username,familyId, dutyId, points, gradeComment
	familyTutor.POST("family/:famId/create/duty", controller.SaveNewDuty())
	familyTutor.GET("/family/:famId/duty/:dutyId/grading", controller.GetGradingByDutyId())
	//familyTutor.POST("/execute/duty", controller.TriggerExecution()) // dutyId
}

func familyMemberRoutes() {
	// has to check if this user is a member of this fam
	familyMember := seapRouter.Group("/api/my/")
	familyMember.Use(controller.FamilyMemberMiddleware())
	familyMember.GET("/family/:famId/duties", controller.GetAllDutiesByFamilyId()) // get all duties associated with the fam id
	familyMember.GET("/family/:famId/duty/:dutyId", controller.GetDutyById())      // get specific duty by id

}

func familyTuteeRoutes() {
	//// execute only if the user has a tutee role in the family
	familyTutee := seapRouter.Group("/api/my/")
	familyTutee.Use(controller.FamilyTuteeMiddleware())

	//// grab gradingId and insert the file into submitted_file table
	//protected.GET("/submit/:gradingId", controller.GetGradingDetailForSubmission())
	//protected.POST("/upload/:gradingId", controller.UploadFilesByTutee)     // username, familyId, dutyId, files
	//protected.POST("/submit/:gradingId/done", controller.SubmitDutyByTutee) // username, familyId, dutyId
}
