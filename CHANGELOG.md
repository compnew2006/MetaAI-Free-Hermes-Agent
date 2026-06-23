# Changelog

All notable changes to this project will be documented in this file.

## [2.0.0] - 2026-06-23

### Added
- Working **cookie-based authentication** (requires only 2 cookies: `datr` and `ecto_1_sess`).
- Full support for **Video Generation** with customizable orientation (`LANDSCAPE`, `VERTICAL`, `SQUARE`) and auto-polling functionality.
- Custom orientations for **Image Generation** (`LANDSCAPE`, `VERTICAL`, `SQUARE`).
- Multi-turn **Conversation Context** support.
- Fully functional REST API server using FastAPI/Uvicorn.
- Auto-refresh scripts for session cookies (`ecto_1_sess`).

### Changed
- Removed deprecated token fetching (`lsd` / `fb_dtsg`) ensuring faster server startup and higher compatibility.

### Fixed
- Proxy support reliability across all API endpoints.

---

## [1.0.0] - 2025-10-12

### Added
- Initial release of the Meta AI Python SDK.
- Basic text generation and image generation.
