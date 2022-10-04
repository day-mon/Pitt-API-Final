# University of Pittsburgh Unoffical REST API

**This is a work in progress.**

This is a REST API for the University of Pittsburgh. It is not affiliated with the University of Pittsburgh in any way.

## Endpoints
- `v1/courses` - Base endpoint for all courses related endpoints
  - POST: `v1/courses/info` - Get information about many courses
    - Request Body (**All non-required fields are optional and will be assumed**):
      - `term` - The term to get courses for. **Required**.
      - `subject` - The subject to get courses for
      - `campus` - The campus to get courses for
      - `enrollment_status` - The enrollment status to get courses for
      - `pageNumber` - The page number the courses are on

        - Code Example in **JavaScript**:
          ```js
          const body = {
            term: '2234',
            subject: 'CS',
            campus: 'UPJ',
            enrollment_status: 'O',
            pageNumber: 1
          }
          
          fetch('https://pittapi.example.com/v1/courses/info', {
            method: 'POST',
            body: JSON.stringify(body),
            headers: {
              'Content-Type': 'application/json'
            }
          })
          ```
    - Response Body:
      - `courses` - Array of course objects that are on that page or page 1 if none specified
  

  