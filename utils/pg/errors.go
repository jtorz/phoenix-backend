// package pg is a wrapper for github.com/lib/pq error
package pg

import (
	"errors"

	"github.com/lib/pq"
	pkgerrors "github.com/pkg/errors"
)

// Error represents an error communicating with the server.
//
// See http://www.postgresql.org/docs/current/static/protocol-error-fields.html for details of the fields
type Error struct {
	Severity         string
	Code             ErrorCode
	Message          string
	Detail           string
	Hint             string
	Position         string
	InternalPosition string
	InternalQuery    string
	Where            string
	Schema           string
	Table            string
	Column           string
	DataTypeName     string
	Constraint       string
	File             string
	Line             string
	Routine          string
}

func (err Error) Error() string {
	if err.Code == NoDataFound {
		return err.Table + " not found"
	}
	return "pq: " + err.Message + err.Table
}

// IsCode checks if the error is caused by the postgres code.
func IsCode(err error, code ErrorCode) bool {
	pgerr := GetPgErr(err)
	if pgerr == nil {
		return false
	}
	return pgerr.Code == code
}

// GetPgErr returns the postgres underlying error if posible.
func GetPgErr(err error) *Error {
	if pgerr, ok := err.(*Error); ok {
		return pgerr
	}

	// try unwrap
	err2 := errors.Unwrap(err)
	if err2 != nil {
		err = err2
	}
	pgErr, ok := pkgerrors.Cause(err).(*pq.Error)
	if !ok {
		return nil
	}
	return &Error{
		Severity:         pgErr.Severity,
		Code:             ErrorCode(pgErr.Code),
		Message:          pgErr.Message,
		Detail:           pgErr.Detail,
		Hint:             pgErr.Hint,
		Position:         pgErr.Position,
		InternalPosition: pgErr.InternalPosition,
		InternalQuery:    pgErr.InternalQuery,
		Where:            pgErr.Where,
		Schema:           pgErr.Schema,
		Table:            pgErr.Table,
		Column:           pgErr.Column,
		DataTypeName:     pgErr.DataTypeName,
		Constraint:       pgErr.Constraint,
		File:             pgErr.File,
		Line:             pgErr.Line,
		Routine:          pgErr.Routine,
	}
}

// ErrorCode is a mapping between the five-character error codes and the
// human readable "condition names". It is derived from the list at
// http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html
type ErrorCode string

const (
	// Class 00 - Successful Completion
	SuccessfulCompletion ErrorCode = "00000"
	// Class 01 - Warning
	Warning                              ErrorCode = "01000"
	WarnDynamicResultSetsReturned        ErrorCode = "0100C"
	WarnImplicitZeroBitPadding           ErrorCode = "01008"
	WarnNullValueEliminatedInSetFunction ErrorCode = "01003"
	WarnPrivilegeNotGranted              ErrorCode = "01007"
	WarnPrivilegeNotRevoked              ErrorCode = "01006"
	WarnStringDataRightTruncation        ErrorCode = "01004"
	WarnDeprecatedFeature                ErrorCode = "01P01"
	// Class 02 - No Data (this is also a warning class per the SQL standard)
	NoData                                ErrorCode = "02000"
	NoAdditionalDynamicResultSetsReturned ErrorCode = "02001"
	// Class 03 - SQL Statement Not Yet Complete
	SQLStatementNotYetComplete ErrorCode = "03000"
	// Class 08 - Connection Exception
	ConnectionException                           ErrorCode = "08000"
	ConnectionDoesNotExist                        ErrorCode = "08003"
	ConnectionFailure                             ErrorCode = "08006"
	SQLclientUnableToEstablishSQLconnection       ErrorCode = "08001"
	SQLserverRejectedEstablishmentOfSQLconnection ErrorCode = "08004"
	TransactionResolutionUnknown                  ErrorCode = "08007"
	ProtocolViolation                             ErrorCode = "08P01"
	// Class 09 - Triggered Action Exception
	TriggeredActionException ErrorCode = "09000"
	// Class 0A - Feature Not Supported
	FeatureNotSupported ErrorCode = "0A000"
	// Class 0B - Invalid Transaction Initiation
	InvalidTransactionInitiation ErrorCode = "0B000"
	// Class 0F - Locator Exception
	LocatorException            ErrorCode = "0F000"
	InvalidLocatorSpecification ErrorCode = "0F001"
	// Class 0L - Invalid Grantor
	InvalidGrantor        ErrorCode = "0L000"
	InvalidGrantOperation ErrorCode = "0LP01"
	// Class 0P - Invalid Role Specification
	InvalidRoleSpecification ErrorCode = "0P000"
	// Class 0Z - Diagnostics Exception
	DiagnosticsException                           ErrorCode = "0Z000"
	StackedDiagnosticsAccessedWithoutActiveHandler ErrorCode = "0Z002"
	// Class 20 - Case Not Found
	CaseNotFound ErrorCode = "20000"
	// Class 21 - Cardinality Violation
	CardinalityViolation ErrorCode = "21000"
	// Class 22 - Data Exception
	DataException                         ErrorCode = "22000"
	ArraySubscriptError                   ErrorCode = "2202E"
	CharacterNotInRepertoire              ErrorCode = "22021"
	DatetimeFieldOverflow                 ErrorCode = "22008"
	DivisionByZero                        ErrorCode = "22012"
	ErrorInAssignment                     ErrorCode = "22005"
	EscapeCharacterConflict               ErrorCode = "2200B"
	IndicatorOverflow                     ErrorCode = "22022"
	IntervalFieldOverflow                 ErrorCode = "22015"
	InvalidArgumentForLogarithm           ErrorCode = "2201E"
	InvalidArgumentForNtileFunction       ErrorCode = "22014"
	InvalidArgumentForNthValueFunction    ErrorCode = "22016"
	InvalidArgumentForPowerFunction       ErrorCode = "2201F"
	InvalidArgumentForWidthBucketFunction ErrorCode = "2201G"
	InvalidCharacterValueForCast          ErrorCode = "22018"
	InvalidDatetimeFormat                 ErrorCode = "22007"
	InvalidEscapeCharacter                ErrorCode = "22019"
	InvalidEscapeOctet                    ErrorCode = "2200D"
	InvalidEscapeSequence                 ErrorCode = "22025"
	NonstandardUseOfEscapeCharacter       ErrorCode = "22P06"
	InvalidIndicatorParameterValue        ErrorCode = "22010"
	InvalidParameterValue                 ErrorCode = "22023"
	InvalidRegularExpression              ErrorCode = "2201B"
	InvalidRowCountInLimitClause          ErrorCode = "2201W"
	InvalidRowCountInResultOffsetClause   ErrorCode = "2201X"
	InvalidTimeZoneDisplacementValue      ErrorCode = "22009"
	InvalidUseOfEscapeCharacter           ErrorCode = "2200C"
	MostSpecificTypeMismatch              ErrorCode = "2200G"
	NullValueNotAllowed                   ErrorCode = "22004"
	NullValueNoIndicatorParameter         ErrorCode = "22002"
	NumericValueOutOfRange                ErrorCode = "22003"
	SequenceGeneratorLimitExceeded        ErrorCode = "2200H"
	StringDataLengthMismatch              ErrorCode = "22026"
	StringDataRightTruncation             ErrorCode = "22001"
	SubstringError                        ErrorCode = "22011"
	TrimError                             ErrorCode = "22027"
	UnterminatedCString                   ErrorCode = "22024"
	ZeroLengthCharacterString             ErrorCode = "2200F"
	FloatingPointException                ErrorCode = "22P01"
	InvalidTextRepresentation             ErrorCode = "22P02"
	InvalidBinaryRepresentation           ErrorCode = "22P03"
	BadCopyFileFormat                     ErrorCode = "22P04"
	UntranslatableCharacter               ErrorCode = "22P05"
	NotAnXMLDocument                      ErrorCode = "2200L"
	InvalidXMLDocument                    ErrorCode = "2200M"
	InvalidXMLContent                     ErrorCode = "2200N"
	InvalidXMLComment                     ErrorCode = "2200S"
	InvalidXMLProcessingInstruction       ErrorCode = "2200T"
	// Class 23 - Integrity Constraint Violation
	IntegrityConstraintViolation ErrorCode = "23000"
	RestrictViolation            ErrorCode = "23001"
	NotNullViolation             ErrorCode = "23502"
	ForeignKeyViolation          ErrorCode = "23503"
	UniqueViolation              ErrorCode = "23505"
	CheckViolation               ErrorCode = "23514"
	ExclusionViolation           ErrorCode = "23P01"
	// Class 24 - Invalid Cursor State
	InvalidCursorState ErrorCode = "24000"
	// Class 25 - Invalid Transaction State
	InvalidTransactionState                         ErrorCode = "25000"
	ActiveSQLTransaction                            ErrorCode = "25001"
	BranchTransactionAlreadyActive                  ErrorCode = "25002"
	HeldCursorRequiresSameIsolationLevel            ErrorCode = "25008"
	InappropriateAccessModeForBranchTransaction     ErrorCode = "25003"
	InappropriateIsolationLevelForBranchTransaction ErrorCode = "25004"
	NoActiveSQLTransactionForBranchTransaction      ErrorCode = "25005"
	ReadOnlySQLTransaction                          ErrorCode = "25006"
	SchemaAndDataStatementMixingNotSupported        ErrorCode = "25007"
	NoActiveSQLTransaction                          ErrorCode = "25P01"
	InFailedSQLTransaction                          ErrorCode = "25P02"
	// Class 26 - Invalid SQL Statement Name
	InvalidSQLStatementName ErrorCode = "26000"
	// Class 27 - Triggered Data Change Violation
	TriggeredDataChangeViolation ErrorCode = "27000"
	// Class 28 - Invalid Authorization Specification
	InvalidAuthorizationSpecification ErrorCode = "28000"
	InvalidPassword                   ErrorCode = "28P01"
	// Class 2B - Dependent Privilege Descriptors Still Exist
	DependentPrivilegeDescriptorsStillExist ErrorCode = "2B000"
	DependentObjectsStillExist              ErrorCode = "2BP01"
	// Class 2D - Invalid Transaction Termination
	InvalidTransactionTermination ErrorCode = "2D000"
	// Class 2F - SQL Routine Exception
	SQLRoutineException                           ErrorCode = "2F000"
	SQLRoutineExFunctionExecutedNoReturnStatement ErrorCode = "2F005"
	SQLRoutineExModifyingSQLDataNotPermitted      ErrorCode = "2F002"
	SQLRoutineExProhibitedSQLStatementAttempted   ErrorCode = "2F003"
	SQLRoutineExReadingSQLDataNotPermitted        ErrorCode = "2F004"
	// Class 34 - Invalid Cursor Name
	InvalidCursorName ErrorCode = "34000"
	// Class 38 - External Routine Exception
	ExternalRoutineException                         ErrorCode = "38000"
	ExternalRoutineExContainingSQLNotPermitted       ErrorCode = "38001"
	ExternalRoutineExModifyingSQLDataNotPermitted    ErrorCode = "38002"
	ExternalRoutineExProhibitedSQLStatementAttempted ErrorCode = "38003"
	ExternalRoutineExReadingSQLDataNotPermitted      ErrorCode = "38004"
	// Class 39 - External Routine Invocation Exception
	ExternalRoutineInvocationException         ErrorCode = "39000"
	ExternalRoutineInExInvalidSQLstateReturned ErrorCode = "39001"
	ExternalRoutineInExNullValueNotAllowed     ErrorCode = "39004"
	ExternalRoutineInExTriggerProtocolViolated ErrorCode = "39P01"
	ExternalRoutineInExSrfProtocolViolated     ErrorCode = "39P02"
	// Class 3B - Savepoint Exception
	SavepointException            ErrorCode = "3B000"
	InvalidSavepointSpecification ErrorCode = "3B001"
	// Class 3D - Invalid Catalog Name
	InvalidCatalogName ErrorCode = "3D000"
	// Class 3F - Invalid Schema Name
	InvalidSchemaName ErrorCode = "3F000"
	// Class 40 - Transaction Rollback
	TransactionRollback                     ErrorCode = "40000"
	TransactionIntegrityConstraintViolation ErrorCode = "40002"
	SerializationFailure                    ErrorCode = "40001"
	StatementCompletionUnknown              ErrorCode = "40003"
	DeadlockDetected                        ErrorCode = "40P01"
	// Class 42 - Syntax Error or Access Rule Violation
	SyntaxErrorOrAccessRuleViolation   ErrorCode = "42000"
	SyntaxError                        ErrorCode = "42601"
	InsufficientPrivilege              ErrorCode = "42501"
	CannotCoerce                       ErrorCode = "42846"
	GroupingError                      ErrorCode = "42803"
	WindowingError                     ErrorCode = "42P20"
	InvalidRecursion                   ErrorCode = "42P19"
	InvalidForeignKey                  ErrorCode = "42830"
	InvalidName                        ErrorCode = "42602"
	NameTooLong                        ErrorCode = "42622"
	ReservedName                       ErrorCode = "42939"
	DatatypeMismatch                   ErrorCode = "42804"
	IndeterminateDatatype              ErrorCode = "42P18"
	CollationMismatch                  ErrorCode = "42P21"
	IndeterminateCollation             ErrorCode = "42P22"
	WrongObjectType                    ErrorCode = "42809"
	UndefinedColumn                    ErrorCode = "42703"
	UndefinedFunction                  ErrorCode = "42883"
	UndefinedTable                     ErrorCode = "42P01"
	UndefinedParameter                 ErrorCode = "42P02"
	UndefinedObject                    ErrorCode = "42704"
	DuplicateColumn                    ErrorCode = "42701"
	DuplicateCursor                    ErrorCode = "42P03"
	DuplicateDatabase                  ErrorCode = "42P04"
	DuplicateFunction                  ErrorCode = "42723"
	DuplicatePreparedStatement         ErrorCode = "42P05"
	DuplicateSchema                    ErrorCode = "42P06"
	DuplicateTable                     ErrorCode = "42P07"
	DuplicateAlias                     ErrorCode = "42712"
	DuplicateObject                    ErrorCode = "42710"
	AmbiguousColumn                    ErrorCode = "42702"
	AmbiguousFunction                  ErrorCode = "42725"
	AmbiguousParameter                 ErrorCode = "42P08"
	AmbiguousAlias                     ErrorCode = "42P09"
	InvalidColumnReference             ErrorCode = "42P10"
	InvalidColumnDefinition            ErrorCode = "42611"
	InvalidCursorDefinition            ErrorCode = "42P11"
	InvalidDatabaseDefinition          ErrorCode = "42P12"
	InvalidFunctionDefinition          ErrorCode = "42P13"
	InvalidPreparedStatementDefinition ErrorCode = "42P14"
	InvalidSchemaDefinition            ErrorCode = "42P15"
	InvalidTableDefinition             ErrorCode = "42P16"
	InvalidObjectDefinition            ErrorCode = "42P17"
	// Class 44 - WITH CHECK OPTION Violation
	WithCheckOptionViolation ErrorCode = "44000"
	// Class 53 - Insufficient Resources
	InsufficientResources      ErrorCode = "53000"
	DiskFull                   ErrorCode = "53100"
	OutOfMemory                ErrorCode = "53200"
	TooManyConnections         ErrorCode = "53300"
	ConfigurationLimitExceeded ErrorCode = "53400"
	// Class 54 - Program Limit Exceeded
	ProgramLimitExceeded ErrorCode = "54000"
	StatementTooComplex  ErrorCode = "54001"
	TooManyColumns       ErrorCode = "54011"
	TooManyArguments     ErrorCode = "54023"
	// Class 55 - Object Not In Prerequisite State
	ObjectNotInPrerequisiteState ErrorCode = "55000"
	ObjectInUse                  ErrorCode = "55006"
	CantChangeRuntimeParam       ErrorCode = "55P02"
	LockNotAvailable             ErrorCode = "55P03"
	// Class 57 - Operator Intervention
	OperatorIntervention ErrorCode = "57000"
	QueryCanceled        ErrorCode = "57014"
	AdminShutdown        ErrorCode = "57P01"
	CrashShutdown        ErrorCode = "57P02"
	CannotConnectNow     ErrorCode = "57P03"
	DatabaseDropped      ErrorCode = "57P04"
	// Class 58 - System Error (errors external to PostgreSQL itself)
	SystemError   ErrorCode = "58000"
	IoError       ErrorCode = "58030"
	UndefinedFile ErrorCode = "58P01"
	DuplicateFile ErrorCode = "58P02"
	// Class F0 - Configuration File Error
	ConfigFileError ErrorCode = "F0000"
	LockFileExists  ErrorCode = "F0001"
	// Class HV - Foreign Data Wrapper Error (SQL/MED)
	FdwError                             ErrorCode = "HV000"
	FdwColumnNameNotFound                ErrorCode = "HV005"
	FdwDynamicParameterValueNeeded       ErrorCode = "HV002"
	FdwFunctionSequenceError             ErrorCode = "HV010"
	FdwInconsistentDescriptorInformation ErrorCode = "HV021"
	FdwInvalidAttributeValue             ErrorCode = "HV024"
	FdwInvalidColumnName                 ErrorCode = "HV007"
	FdwInvalidColumnNumber               ErrorCode = "HV008"
	FdwInvalidDataType                   ErrorCode = "HV004"
	FdwInvalidDataTypeDescriptors        ErrorCode = "HV006"
	FdwInvalidDescriptorFieldIdentifier  ErrorCode = "HV091"
	FdwInvalidHandle                     ErrorCode = "HV00B"
	FdwInvalidOptionIndex                ErrorCode = "HV00C"
	FdwInvalidOptionName                 ErrorCode = "HV00D"
	FdwInvalidStringLengthOrBufferLength ErrorCode = "HV090"
	FdwInvalidStringFormat               ErrorCode = "HV00A"
	FdwInvalidUseOfNullPointer           ErrorCode = "HV009"
	FdwTooManyHandles                    ErrorCode = "HV014"
	FdwOutOfMemory                       ErrorCode = "HV001"
	FdwNoSchemas                         ErrorCode = "HV00P"
	FdwOptionNameNotFound                ErrorCode = "HV00J"
	FdwReplyHandle                       ErrorCode = "HV00K"
	FdwSchemaNotFound                    ErrorCode = "HV00Q"
	FdwTableNotFound                     ErrorCode = "HV00R"
	FdwUnableToCreateExecution           ErrorCode = "HV00L"
	FdwUnableToCreateReply               ErrorCode = "HV00M"
	FdwUnableToEstablishConnection       ErrorCode = "HV00N"
	// Class P0 - PL/pgSQL Error
	PlpgsqlError   ErrorCode = "P0000"
	RaiseException ErrorCode = "P0001"
	NoDataFound    ErrorCode = "P0002"
	TooManyRows    ErrorCode = "P0003"
	// Class XX - Internal Error
	InternalError  ErrorCode = "XX000"
	DataCorrupted  ErrorCode = "XX001"
	IndexCorrupted ErrorCode = "XX002"
)
