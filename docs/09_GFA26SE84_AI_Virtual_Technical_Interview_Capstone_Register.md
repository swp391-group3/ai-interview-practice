# CAPSTONE PROJECT REGISTER

**Class**: **Duration time**: From ... 11/05/2026 ... to ... /08/2026 ...

**(\*) Profession:** Software Engineer **Specialty:** Golang - NextJS - AI Integration

**(\*) Kinds of person make registers:** Lecturer Students !

## Register information for the supervisor

| **No.**      | **Full name**    | **Phone**    | **E-Mail**                                        | **Title** |
| ------------ | ---------------- | ------------ | ------------------------------------------------- | --------- |
| Supervisor 1 | Nguyễn Thế Hoàng | 0986 628 525 | [hoangnt20@fe.edu.vn](mailto:hoangnt20@fe.edu.vn) | Mr.       |

## Register information for students

| **No.** | **Full name**                | **Student code** | **Phone** | **Email** | **Role in Group**               |
| ------- | ---------------------------- | ---------------- | --------- | --------- | ------------------------------- |
| 1       | Tô Chí Bảo                   | SE190084         |           |           | Full Stack                      |
| 2       | Huỳnh Minh Khang             | SE192197         |           |           | Full Stack                      |
| 3       | Nguyễn Huỳnh Nhật<br><br>Anh | SE190291         |           |           | Team Leader /<br><br>Full Stack |
| 4       | Nguyễn Tấn Trọng             | SE190353         |           |           | Full Stack                      |
| 5       | Đặng Phương Nam              | SE192107         |           |           | Full Stack                      |

## Register content of the Capstone Project

### (\*) 3.1. Capstone Project name

-   - **English:** Design and Development of an AI-Powered Technical Interview Simulation Platform with a 3D Virtual Interviewer
        - **Vietnamese:** Xây dựng nền tảng mô phỏng phỏng vấn kỹ thuật trực tuyến ứng dụng trí tuệ nhân tạo và nhân vật 3D ảo

#### Context

Technical interview preparation is an important step for students and software developers who are seeking internship or employment opportunities. However, current preparation methods often rely on generic question banks, self-study materials, or occasional mock interviews. These approaches do not sufficiently reflect the personalized and adaptive nature of real technical interviews. The main challenges include:

-   - **Lack of personalization:** Most practice platforms provide general questions without considering the specific requirements of the Job Description (JD) the candidate is actually applying for.
        - **Limited access to realistic mock interviews:** Candidates may not always have an experienced interviewer available for repeated one-on-one practice, making it difficult to build confidence and interview readiness.
        - **Static and non-adaptive questioning:** Traditional question banks usually present predefined questions and cannot ask contextual follow-up questions based on the candidate's previous answers or depth of knowledge.
        - **Insufficient evaluation and actionable feedback:** Candidates often receive only a final score or no structured assessment at all, which makes it difficult to identify technical strengths, knowledge gaps, and specific areas for improvement relative to the target job.

#### Proposed Solutions

To address these limitations, the project proposes an AI-Powered Virtual Technical Interview Simulation Platform. The system analyzes a specific Job Description (JD) provided by the candidate, generates a personalized interview strategy tailored to those exact job requirements, generates a personalized interview strategy, conducts a real-time voice-based interview through a 3D virtual interviewer, and provides a detailed AI-generated performance report. The proposed solution focuses on four main pillars:

#### Job Description (JD) Requirements Analysis

The system allows candidates to input or upload a specific Job Description (JD) and automatically extracts relevant technical requirements such as programming languages, frameworks, databases, tools, and domain knowledge. The extracted information is used to construct an interview blueprint that strictly aligns with the employer's expectations.

#### JD-Driven Personalized Interview Planning

Based on the extracted requirements from the JD, the AI determines the optimal interview topics, depth, and question difficulty. This ensures that the candidate is evaluated on the exact technical stack and problem-solving scenarios required by the specific role they are targeting.

#### Voice-Based 3D Virtual AI Interviewer

The system provides a one-on-one interview experience in which candidates communicate with an AI interviewer using voice. Speech recognition and speech synthesis services support natural conversation, while a 3D virtual interviewer uses synchronized speaking and lip-sync animations. The AI maintains session context and asks adaptive follow-up questions based on previous answers and the core requirements of the JD.

#### AI Evaluation & Detailed Performance Reporting

After the interview, the system evaluates the candidate's technical accuracy, depth of understanding, problem-solving ability, answer relevance, and communication clarity against the JD's requirements. A detailed report presents the overall score, technical-domain scores, question-level feedback, strengths, weaknesses, and personalized improvement recommendations to help the candidate secure the specific job.

### Functional requirements

#### Admin

**Account Management**

- View registered candidate accounts and account details
- Lock or unlock candidate accounts when necessary.

#### Interview Session Management

- View and filter interview sessions by candidate, status or date.
- View interview session details and track session statuses, including completed, failed, or interrupted sessions.

#### Technical Domain Management

- Manage technical domains and technologies used as interview context.

#### Interview Configuration Management

- Configure interview difficulty levels, duration constraints, evaluation criteria, and interview behavior guidelines.
- Configure AI interview behavior and evaluation settings.

#### Virtual Interviewer Avatar & Voice Management

- Manage available 3D virtual interviewer avatars.
- Manage supported interviewer voice profiles used during voice-based interview sessions.

#### System Analytics & Reporting

- View statistics for total, completed, failed, and interrupted interview sessions.
- Analyze interview volume and average candidate performance.
- View commonly identified technical weaknesses across completed interviews.

#### Subscription & Revenue Management

- Manage pricing plans, subscription tiers, or interview credit packages.
- View user transaction history and payment statuses.
- Generate and view revenue reports by day, month, or year.

#### Candidate

**Authentication & Profile Management**

- Register, log in, log out.
- View and update personal profile information and account credentials.

#### Job Description (JD) Management

- Upload, paste, or input Job Descriptions (JD) used for interview preparation.
- Allow the system to extract required technical skills, technologies, and domain knowledge from the JD content.
- Review and edit the extracted requirements before starting an interview.

#### Interview Configuration

- Select a previously analyzed JD or input a new one for the interview.
- Select an analyzed JD and configure interview parameters such as difficulty, duration, and number of questions before starting the interview.
- Start a personalized technical interview based strictly on the requirements of the chosen JD.

#### AI Virtual Interview

- Join a one-on-one technical interview session with a 3D AI virtual interviewer.
- Communicate with the interviewer using voice and hear questions through synthesized speech.
- Experience synchronized speaking and lip-sync animation from the 3D interviewer.
- Receive adaptive technical questions and contextual follow-up questions based on the JD requirements and previous answers.

#### Interview History

- View previous interview sessions, including target JD, interview date, status, and overall score.
- Revisit completed interview results and performance reports.

#### AI Evaluation & Performance Report

- View an overall interview performance score and scores by technical domain.
- Review competency assessments covering technical accuracy, depth of understanding, problem-solving, answer relevance, and communication clarity against the JD.
- View question-level feedback, identified strengths and weaknesses, and personalized improvement recommendations.

#### Billing & Payment Management

- View current subscription plan and remaining interview credits.
- Upgrade to a Premium plan or purchase additional interview credits via a Payment Gateway.
- View personal billing and transaction history.

### Non-Functional requirements

#### Performance

- Common operations such as login, profile viewing, interview history retrieval, and report viewing should respond within **3 seconds** under normal usage.
- JD analysis and AI evaluation should provide a loading or progress indication while processing.
- The system should support at least **20 concurrent users** during normal testing.
- The interview interface should maintain acceptable responsiveness during voice communication and virtual interviewer animation.

#### Security

- Require authentication for protected functions.
- Enforce role-based access control between **Admin and Candidate**.
- Store user passwords using secure password hashing.
- Use secure authentication mechanisms and HTTPS for client-server communication.
- Protect against common web vulnerabilities such as **SQL Injection, XSS, and unauthorized access**.
- Candidates must only be able to access their own profile, interview sessions, results, and transaction records.

#### Reliability & Data Integrity

- Interview session data should be saved correctly when an interview is completed.
- The system should correctly record interview statuses such as **completed, failed, and interrupted**.
- Interview results should be associated with the correct candidate, JD, and interview session.
- Interview credits and payment records should remain consistent when a purchase is completed.
- The system should handle temporary failures from external AI, speech, or payment services by displaying an appropriate error message.

#### Usability

- Provide a simple and consistent web interface suitable for first-time users.
- Provide clear validation and error messages for invalid input and failed operations.
- Clearly display interview progress, session status, scores, and performance feedback.
- Provide instructions or guidance for microphone and voice-based interview functionality.
- The web application should be responsive on commonly used **desktop and mobile browsers**.

### (\*) 3.2. Main proposal content (including result and product)

#### A Theory and practice

-   -   - Students should apply the software development process and UML 2.0 in system analysis and modeling.
            - The documents include User Requirements, Software Requirement Specification (SRS), Architecture Design, Detail Design, System Implementation, Testing Document, Installation Guide, source code, and deployable software packages.
            - Source Code Management: GitHub.

#### Server-side technologies

-   -   -   - Server: Go (Golang). - Database: PostgreSQL. - Security: JWT-based authentication and role-based authorization. - AI Integration: Large Language Model APIs for JD analysis, technical profile extraction, personalized interview generation, adaptive follow-up questioning, and answer evaluation. - Voice Interaction: Speech-to-Text and Text-to-Speech services for real-time interview conversation. - Payment Integration: Integration with third-party payment gateways (e.g., VNPay, MoMo, or Stripe) for transaction processing.

#### Client-side technologies

-   -   -   - Web Client: Next.js with web-based 3D avatar rendering, synchronized speaking animation, and lip-sync integration.

#### B. Products

- **Web application:** A deployed full-stack web platform supporting two roles: **Candidate** (JD analysis, subscription management, payment processing, requirements extraction, voice-based interviews with a 3D AI interviewer based on JD, interview history, and detailed AI performance reports) and **Admin** (account management, revenue monitoring, interview session monitoring, interview configuration, target position and technical domain management, virtual interviewer avatar and voice management, and system analytics).

#### C. Proposed Tasks

- Task package 1: Design and deploy the PostgreSQL database and core Golang back-end services.
- Task package 2: Develop JD analysis, technical requirement extraction, and AI interview orchestration.
- Task package 3: Develop the Candidate web application, voice-based interview workflow, 3D virtual interviewer, interview history, detailed performance reports, and payment gateway integration.
- Task package 4: Develop the Admin web application, system analytics, testing, documentation, and real deployment.

| **Supervisor**<br><br>(Sign and full name)<br><br>Nguyễn Thế Hoàng | HCM, date ... / ... /2026<br><br>**On behalf of the Registers**<br><br>(Sign and full name)<br><br>Nguyễn Huỳnh Nhật Anh |
| ------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
