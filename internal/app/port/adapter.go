package port

type Adapter struct {
	RestApiAdapter RestApiAdapter
}
type RestApiAdapter interface {
	// MasterDataGet(in domain.MasterDataApiInput) (*domain.MasterDataApiOutput, error)
	// MasterDataProvinceGet(in domain.MasterDataProvinceApiInput) (*domain.MasterDataProvinceApiOutput, error)
	// MasterDataDistrictGet(in domain.MasterDataDistrictApiInput) (*domain.MasterDataDistrictApiOutput, error)
	// MasterDataSubDistrictGet(in domain.MasterDataSubDistrictApiInput) (*domain.MasterDataSubDistrictApiOutput, error)
	// LearnerProfileGet(in domain.LearnerProfileApiInput) (*domain.ProfileGetByLearnerIdApiOutput, error)

	// GetMentorProfile(ctx context.Context, mentorId string) (domain.MentorGetProfileOutput, error)
	// GetCourseDetail(ctx context.Context, courseId string) (*domain.CourseDetailOutput, error)
	// TrainingGetByIds(in domain.TrainingGetByIdsInput) (*domain.TrainingGetByIdsOutput, error)
	// CourseGetByIds(in domain.CourseGetByIdsInput) (*domain.CourseGetByIdsOutput, error)

	// GetAcademyMemberInfomationByIdAndAcademyId(ctx context.Context, memberId string, academyId string) (*domain.AcademyMemberGetInfomationByIdAndAcademyIdOutput, error)
	// GetLearnerCourseProgress(ctx context.Context, learnerId string, courseIds []string) ([]domain.LearnerCourseProgress, error)
	// GetLearnerTrainingProgress(ctx context.Context, learnerId string, trainingIds []string) ([]domain.LearnerTrainingProgress, error)
	// LearnerSkillUpdate(ctx context.Context, learnerId string, req payload.LearnerSkillUpdateRequest) error
	// LearnerInterestGet(in domain.LearnerInterestApiInput) (*domain.LearnerInterestApiOutput, error)
	// FileUpdate(ctx context.Context, in domain.FileUpdateInput) (*domain.FileUpdateResponse, error)
	// FileLink(ctx context.Context, in domain.FileLinkInput) error
	// FileUnlink(ctx context.Context, in domain.FileUnlinkInput) error
}
