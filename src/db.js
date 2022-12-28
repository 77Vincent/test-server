import {migrate} from "postgres-migrations"

export default async function () {
    const dbConfig = {
        database: "main",
        user: "dev",
        password: "dev",
        host: "localhost",
        port: 5444,

        // Default: false for backwards-compatibility
        // This might change!
        ensureDatabaseExists: true,

        // Default: "postgres"
        // Used when checking/creating "database-name"
        defaultDatabase: "postgres"
    }

    await migrate(dbConfig, "./config/db/migrations")
}