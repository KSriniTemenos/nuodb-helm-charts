package testlib

const CA_CERT_FILE = "ca.cert"
const CA_CERT_SECRET = "nuodb-ca-cert"

const KEYSTORE_FILE = "nuoadmin.p12"
const KEYSTORE_SECRET = "nuodb-keystore"

const TRUSTSTORE_FILE = "nuoadmin-truststore.p12"
const TRUSTSTORE_SECRET = "nuodb-truststore"

const NUOCMD_FILE = "nuocmd.pem"
const NUOCMD_SECRET = "nuodb-client-pem"

const SECRET_PASSWORD = "changeIt"

const TEARDOWN_ADMIN = "admin"
const TEARDOWN_BACKUP = "backup"
const TEARDOWN_DATABASE = "database"
const TEARDOWN_RESTORE = "database"
const TEARDOWN_SECRETS = "secrets"


const ADMIN_HELM_CHART_PATH = "../../stable/admin"
const BACKUP_HELM_CHART_PATH = "../../stable/backup"
const DATABASE_HELM_CHART_PATH = "../../stable/database"
const RESTORE_HELM_CHART_PATH = "../../stable/restore"
const THP_HELM_CHART_PATH = "../../stable/transparent-hugepage"


const LAST_BACKUP_PREFIX string = "nuodb-backup/last_created"
const IMPORT_ARCHIVE_URL = "http://download.nuohub.org/ce_releases/restore.bak.tz"