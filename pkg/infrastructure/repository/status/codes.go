package status

type RepositoryInternalStatus string

const (
	Success                        RepositoryInternalStatus = "OK"
	InternalServerError            RepositoryInternalStatus = "INTERNAL_SERVER_ERROR"
	EntityWithSameKeyAlreadyExists RepositoryInternalStatus = "ENTITY_WITH_SAME_KEY_ALREADY_EXISTS"
	OverlappingPeriodDates         RepositoryInternalStatus = "OVERLAPPING_PERIOD_DATES"
)
