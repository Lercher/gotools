package main

import (
	"fmt"
	"log"

	"github.com/ajitpratap0/GoSQLX/pkg/gosqlx"
	"github.com/ajitpratap0/GoSQLX/pkg/sql/keywords"

    "github.com/ha1tch/tsqlparser"
)

func main() {
    input := `
    BEGIN TRANSACTION;
        DECLARE @myVar INT = 10;
        IF @myVar > 5 BEGIN
            SELECT 
              a = A
            , b as [B plus]
            , c = @myVar 
            , d = NULL
            , e = (select top 1 name from dbo.users where active = 1)
            FROM dbo.users WITH (NOLOCK) WHERE active = true
            
            ;WITH cte AS (
                SELECT id, name FROM dbo.users WHERE active = 1
            ), cte2 AS (
                SELECT id, name FROM dbo.users WHERE active = 2
            )
            SELECT o.*, cte.*, cte2.name as cte2_name
            FROM cte
            left outer join dbo.orders o on cte.id = o.user_id
            left join cte2 on cte.id = cte2.id
            ;
        END ELSE BEGIN
            INSERT INTO users (name, email) VALUES ('Jane Roe', 'jane.roe@example.com'), ('Jamie Roe', 'jamie.roe@example.com');
        END
        RAISERROR('This is test message #7 severity 0', 0, 7) WITH NOWAIT;
        select 1 from dbo.TableValuedFunction('p1', 2, -5);
    ROLLBACK;
    `
    
    program, errors := tsqlparser.Parse(input)
    if len(errors) > 0 {
        for _, err := range errors {
            fmt.Println("Error:", err)
        }
        return
    }
    
    for i, stmt := range program.Statements {
        fmt.Printf("Statement %d: %s\n", i+1, stmt.String())
    }
}

func main2() {
	// Parse SQL in one line - that's it!
	ast, err := gosqlx.ParseWithDialect(`
    -- BEGIN TRANSACTION;
        -- DECLARE @myVar INT = 10;
        -- IF @myVar > 5 BEGIN
            SELECT 
              a = A
            , b as [B plus]
            , c = @myVar 
            , d = NULL
            , e = (select top 1 name from dbo.users where active = 1)
            FROM dbo.users WITH (NOLOCK) WHERE active = true
            
            ;WITH cte AS (
                SELECT id, name FROM dbo.users WHERE active = 1
            ), cte2 AS (
                SELECT id, name FROM dbo.users WHERE active = 2
            )
            SELECT o.*, cte.*, cte2.name as cte2_name
            FROM cte
            left outer join dbo.orders o on cte.id = o.user_id
            left join cte2 on cte.id = cte2.id
            ;
        -- END ELSE BEGIN
            INSERT INTO users (name, email) VALUES ('Jane Roe', 'jane.roe@example.com'), ('Jamie Roe', 'jamie.roe@example.com');
        -- END
        -- RAISERROR('This is test message #7 severity 0', 0, 7) WITH NOWAIT;
        -- select 1 from dbo.TableValuedFunction('p1', 2, -5);
    -- ROLLBACK;
    `, keywords.DialectSQLServer)
	if err != nil {
		log.Fatal(err)
	}

	// Success!
	fmt.Printf("Successfully parsed SQL!\n")
	fmt.Printf("Type: %T\n", ast)
	fmt.Printf("Statements: %d\n", len(ast.Statements))
	for i, stmt := range ast.Statements {
		fmt.Printf("Statement %d: %v\n", i+1, stmt.TokenLiteral())
	}
}
