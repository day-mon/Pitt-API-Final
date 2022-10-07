# University of Pittsburgh Unoffical REST API

**This is a work in progress.**

This is a REST API for the University of Pittsburgh. It is not affiliated with the University of Pittsburgh in any way.

## Endpoints
- `/v1/api/courses` - Base endpoint for all courses related endpoints
  - POST: `/courses/info` - Get information about many courses
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
          
          fetch('https://pittapi.example.com/v1/api/courses/info', {
            method: 'POST',
            body: JSON.stringify(body),
            headers: {
              'Content-Type': 'application/json'
            }
          })
          ```
    - Response Body:
      - `courses` - Array of course objects that are on that page or page 1 if none specified
- `/v1/api/course` - Base endpoint for all courses related endpoints
  - GET `/:term/:number` - Get information about a course
    - Get Parameters:
      - `term` - The term to get the course for
      - `number` - The course number to get the course for
        - Code Example in **JavaScript**:
          ```js
          fetch('https://pittapi.example.com/v1/api/course/2234/12345')
          ```
    - Response Body:
      - `course` - The course object

- `/v1/api/laundry`
  -  GET `/:dormitory` - Get information about all laundry machines in a dormitory
    -  Get Parameters:
      - `dormitory` - The dormitory to get laundry information for
        - Code Example in **JavaScript**:
          ```js
          fetch('https://pittapi.example.com/v1/api/laundry/willow')
          ```
        - Response Body:
          - `machines` - Array of machine objects