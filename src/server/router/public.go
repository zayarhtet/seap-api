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
}

func individualRoutes() {
	protected := seapRouter.Group("/api/my/")
	protected.Use(controller.IndividualMiddleware())
	protected.GET("/valid", controller.Welcome())
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
	familyTutor.POST("family/:famId/create/grade", controller.AddNewGrade())        // username,familyId, dutyId, points, gradeComment
	familyTutor.POST("family/:famId/create/duty", controller.SaveNewDuty())
	familyTutor.POST("/cdn/upload/:dutyId/given-file", controller.SaveGivenFiles())
	familyTutor.GET("/family/:famId/duty/:dutyId/grading", controller.GetGradingByDutyId())

	familyTutor.DELETE("/family/:famId", controller.DeleteFamily())
	familyTutor.DELETE("/family/:famId/duty/:dutyId", controller.DeleteDuty())
	//familyTutor.POST("/execute/duty", controller.TriggerExecution()) // dutyId
}

func familyMemberRoutes() {
	// has to check if this user is a member of this fam
	familyMember := seapRouter.Group("/api/my/")
	familyMember.Use(controller.FamilyMemberMiddleware())
	familyMember.GET("/family/:famId/duties", controller.GetAllDutiesByFamilyId())               // get all duties associated with the fam id
	familyMember.GET("/family/:famId/duty/:dutyId", controller.GetDutyById())                    // get specific duty by id
	familyMember.GET("cdn/download/:famId/:dutyId/file/:fileId", controller.DownloadGivenFile()) // get specific duty by id
	familyMember.GET("/family/:famId/myrole", controller.GetMyRoleInFamily())
	familyMember.GET("cdn/download/:famId/family-icon", controller.CDNProfileImage())
	familyMember.GET("cdn/download/family/:famId/duty/:dutyId/submitted-file/:fileId", controller.DownloadSubmittedFile())

}

func familyTuteeRoutes() {
	//// execute only if the user has a tutee role in the family
	familyTutee := seapRouter.Group("/api/my/")
	familyTutee.Use(controller.FamilyTuteeMiddleware())

	familyTutee.GET("family/:famId/duty/:dutyId/my-grading", controller.GetMyGradingDetail())
	familyTutee.POST("cdn/upload/family/:famId/duty/:dutyId/submitted-file", controller.UploadSubmittedFiles()) // username, familyId, dutyId, files
	familyTutee.DELETE("cdn/delete/family/:famId/duty/:dutyId/submitted-file/:fileId", controller.DeleteSubmittedFile())
	familyTutee.POST("/family/:famId/duty/:dutyId/submit/:gradingId/done", controller.SubmitDutyByTutee()) // username, familyId, dutyId
}
