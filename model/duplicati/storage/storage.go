package storage

import "github.com/go-sql-driver/mysql"

// DB Report defines the model.
type DBReport struct {
	ID                                                   uint32         `db:"id"`
	Server                                               string         `db:"Server"`
	DataDeletedFiles                                     uint32         `db:"Data_DeletedFiles"`
	DataDeletedFolders                                   uint32         `db:"Data_DeletedFolders"`
	DataModifiedFiles                                    uint32         `db:"Data_ModifiedFiles"`
	DataExaminedFiles                                    uint32         `db:"Data_ExaminedFiles"`
	DataOpenedFiles                                      uint32         `db:"Data_OpenedFiles"`
	DataAddedFiles                                       uint32         `db:"Data_AddedFiles"`
	DataSizeOfModifiedFiles                              uint32         `db:"Data_SizeOfModifiedFiles"`
	DataSizeOfAddedFiles                                 uint32         `db:"Data_SizeOfAddedFiles"`
	DataSizeOfExaminedFiles                              uint32         `db:"Data_SizeOfExaminedFiles"`
	DataSizeOfOpenedFiles                                uint32         `db:"Data_SizeOfOpenedFiles"`
	DataNotProcessedFiles                                uint32         `db:"Data_NotProcessedFiles"`
	DataAddedFolders                                     uint32         `db:"Data_AddedFolders"`
	DataTooLargeFiles                                    uint32         `db:"Data_TooLargeFiles"`
	DataFilesWithError                                   uint32         `db:"Data_FilesWithError"`
	DataModifiedFolders                                  uint32         `db:"Data_ModifiedFolders"`
	DataModifiedSymlinks                                 uint32         `db:"Data_ModifiedSymlinks"`
	DataAddedSymlinks                                    uint32         `db:"Data_AddedSymlinks"`
	DataDeletedSymlinks                                  uint32         `db:"Data_DeletedSymlinks"`
	DataPartialBackup                                    string         `db:"Data_PartialBackup"`
	DataDryrun                                           string         `db:"Data_Dryrun"`
	DataMainOperation                                    string         `db:"Data_MainOperation"`
	DataCompactResults                                   uint32         `db:"Data_CompactResults"`
	DataVacuumResults                                    uint32         `db:"Data_VacuumResults"`
	DataDeleteResults                                    uint32         `db:"Data_DeleteResults"`
	DataRepairResults                                    uint32         `db:"Data_RepairResults"`
	DataTestResultsMainOperation                         string         `db:"Data_TestResults_MainOperation"`
	DataTestResultsVerificationsActualLength             uint32         `db:"Data_TestResults_VerificationsActualLength"`
	DataTestResultsVerificationsKey                      string         `db:"Data_TestResults_Verifications_Key"`
	DataTestResultsParsedResult                          string         `db:"Data_TestResults_ParsedResult"`
	DataTestResultsVersion                               string         `db:"Data_TestResults_Version"`
	DataTestResultsEndTime                               string         `db:"Data_TestResults_EndTime"`
	DataTestResultsBeginTime                             string         `db:"Data_TestResults_BeginTime"`
	DataTestResultsDuration                              string         `db:"Data_TestResults_Duration"`
	DataTestResultsMessagesActualLength                  uint32         `db:"Data_TestResults_MessagesActualLength"`
	DataTestResultsWarningsActualLength                  uint32         `db:"Data_TestResults_WarningsActualLength"`
	DataTestResultsErrorsActualLength                    uint32         `db:"Data_TestResults_ErrorsActualLength"`
	DataTestResultsBackendStatisticsRemoteCalls          uint32         `db:"Data_TestResults_BackendStatistics_RemoteCalls"`
	DataTestResultsBackendStatisticsBytesUploaded        uint32         `db:"Data_TestResults_BackendStatistics_BytesUploaded"`
	DataTestResultsBackendStatisticsBytesDownloaded      uint32         `db:"Data_TestResults_BackendStatistics_BytesDownloaded"`
	DataTestResultsBackendStatisticsFilesUploaded        uint32         `db:"Data_TestResults_BackendStatistics_FilesUploaded"`
	DataTestResultsBackendStatisticsFilesDownloaded      uint32         `db:"Data_TestResults_BackendStatistics_FilesDownloaded"`
	DataTestResultsBackendStatisticsFilesDeleted         uint32         `db:"Data_TestResults_BackendStatistics_FilesDeleted"`
	DataTestResultsBackendStatisticsFoldersCreated       uint32         `db:"Data_TestResults_BackendStatistics_FoldersCreated"`
	DataTestResultsBackendStatisticsRetryAttempts        uint32         `db:"Data_TestResults_BackendStatistics_RetryAttempts"`
	DataTestResultsBackendStatisticsUnknownFileSize      uint32         `db:"Data_TestResults_BackendStatistics_UnknownFileSize"`
	DataTestResultsBackendStatisticsUnknownFileCount     uint32         `db:"Data_TestResults_BackendStatistics_UnknownFileCount"`
	DataTestResultsBackendStatisticsKnownFileCount       uint32         `db:"Data_TestResults_BackendStatistics_KnownFileCount"`
	DataTestResultsBackendStatisticsKnownFileSize        uint32         `db:"Data_TestResults_BackendStatistics_KnownFileSize"`
	DataTestResultsBackendStatisticsLastBackupDate       string         `db:"Data_TestResults_BackendStatistics_LastBackupDate"`
	DataTestResultsBackendStatisticsBackupListCount      uint32         `db:"Data_TestResults_BackendStatistics_BackupListCount"`
	DataTestResultsBackendStatisticsTotalQuotaSpace      uint32         `db:"Data_TestResults_BackendStatistics_TotalQuotaSpace"`
	DataTestResultsBackendStatisticsFreeQuotaSpace       uint32         `db:"Data_TestResults_BackendStatistics_FreeQuotaSpace"`
	DataTestResultsBackendStatisticsAssignedQuotaSpace   uint32         `db:"Data_TestResults_BackendStatistics_AssignedQuotaSpace"`
	DataTestResultsBackendStatisticsReportedQuotaError   string         `db:"Data_TestResults_BackendStatistics_ReportedQuotaError"`
	DataTestResultsBackendStatisticsReportedQuotaWarning string         `db:"Data_TestResults_BackendStatistics_ReportedQuotaWarning"`
	DataTestResultsBackendStatisticsMainOperation        string         `db:"Data_TestResults_BackendStatistics_MainOperation"`
	DataTestResultsBackendStatisticsParsedResult         string         `db:"Data_TestResults_BackendStatistics_ParsedResult"`
	DataTestResultsBackendStatisticsVersion              string         `db:"Data_TestResults_BackendStatistics_Version"`
	DataTestResultsBackendStatisticsEndTime              string         `db:"Data_TestResults_BackendStatistics_EndTime"`
	DataTestResultsBackendStatisticsBeginTime            string         `db:"Data_TestResults_BackendStatistics_BeginTime"`
	DataTestResultsBackendStatisticsDuration             string         `db:"Data_TestResults_BackendStatistics_Duration"`
	DataTestResultsBackendStatisticsMessagesActualLength uint32         `db:"Data_TestResults_BackendStatistics_MessagesActualLength"`
	DataTestResultsBackendStatisticsWarningsActualLength uint32         `db:"Data_TestResults_BackendStatistics_WarningsActualLength"`
	DataTestResultsBackendStatisticsErrorsActualLength   uint32         `db:"Data_TestResults_BackendStatistics_ErrorsActualLength"`
	DataParsedResult                                     string         `db:"Data_ParsedResult"`
	DataVersion                                          string         `db:"Data_Version"`
	DataEndTime                                          string         `db:"Data_EndTime"`
	DataBeginTime                                        string         `db:"Data_BeginTime"`
	DataDuration                                         string         `db:"Data_Duration"`
	DataMessagesActualLength                             uint32         `db:"Data_MessagesActualLength"`
	DataWarningsActualLength                             uint32         `db:"Data_WarningsActualLength"`
	DataErrorsActualLength                               uint32         `db:"Data_ErrorsActualLength"`
	DataBackendStatisticsRemoteCalls                     uint32         `db:"Data_BackendStatistics_RemoteCalls"`
	DataBackendStatisticsBytesUploaded                   uint32         `db:"Data_BackendStatistics_BytesUploaded"`
	DataBackendStatisticsBytesDownloaded                 uint32         `db:"Data_BackendStatistics_BytesDownloaded"`
	DataBackendStatisticsFilesUploaded                   uint32         `db:"Data_BackendStatistics_FilesUploaded"`
	DataBackendStatisticsFilesDownloaded                 uint32         `db:"Data_BackendStatistics_FilesDownloaded"`
	DataBackendStatisticsFilesDeleted                    uint32         `db:"Data_BackendStatistics_FilesDeleted"`
	DataBackendStatisticsFoldersCreated                  uint32         `db:"Data_BackendStatistics_FoldersCreated"`
	DataBackendStatisticsRetryAttempts                   uint32         `db:"Data_BackendStatistics_RetryAttempts"`
	DataBackendStatisticsUnknownFileSize                 uint32         `db:"Data_BackendStatistics_UnknownFileSize"`
	DataBackendStatisticsUnknownFileCount                uint32         `db:"Data_BackendStatistics_UnknownFileCount"`
	DataBackendStatisticsKnownFileCount                  uint32         `db:"Data_BackendStatistics_KnownFileCount"`
	DataBackendStatisticsKnownFileSize                   uint32         `db:"Data_BackendStatistics_KnownFileSize"`
	DataBackendStatisticsLastBackupDate                  string         `db:"Data_BackendStatistics_LastBackupDate"`
	DataBackendStatisticsBackupListCount                 uint32         `db:"Data_BackendStatistics_BackupListCount"`
	DataBackendStatisticsTotalQuotaSpace                 uint32         `db:"Data_BackendStatistics_TotalQuotaSpace"`
	DataBackendStatisticsFreeQuotaSpace                  uint32         `db:"Data_BackendStatistics_FreeQuotaSpace"`
	DataBackendStatisticsAssignedQuotaSpace              uint32         `db:"Data_BackendStatistics_AssignedQuotaSpace"`
	DataBackendStatisticsReportedQuotaError              string         `db:"Data_BackendStatistics_ReportedQuotaError"`
	DataBackendStatisticsReportedQuotaWarning            string         `db:"Data_BackendStatistics_ReportedQuotaWarning"`
	DataBackendStatisticsMainOperation                   string         `db:"Data_BackendStatistics_MainOperation"`
	DataBackendStatisticsParsedResult                    string         `db:"Data_BackendStatistics_ParsedResult"`
	DataBackendStatisticsVersion                         string         `db:"Data_BackendStatistics_Version"`
	DataBackendStatisticsEndTime                         string         `db:"Data_BackendStatistics_EndTime"`
	DataBackendStatisticsBeginTime                       string         `db:"Data_BackendStatistics_BeginTime"`
	DataBackendStatisticsDuration                        string         `db:"Data_BackendStatistics_Duration"`
	DataBackendStatisticsMessagesActualLength            uint32         `db:"Data_BackendStatistics_MessagesActualLength"`
	DataBackendStatisticsWarningsActualLength            uint32         `db:"Data_BackendStatistics_WarningsActualLength"`
	DataBackendStatisticsErrorsActualLength              uint32         `db:"Data_BackendStatistics_ErrorsActualLength"`
	ExtraOperationName                                   string         `db:"Extra_OperationName"`
	Extrabackupname                                      string         `db:"Extra_backup_name"`
	CreatedAt                                            mysql.NullTime `db:"created_at"`
	UpdatedAt                                            mysql.NullTime `db:"updated_at"`
	DeletedAt                                            mysql.NullTime `db:"deleted_at"`
}

/*
func AddReport(db Connection, report JsonReport) bool {

	status := true

	return status
}

*/
