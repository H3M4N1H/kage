#!/bin/sh
# Enable colors for remote viewers even if the launching environment sets NO_COLOR.
cd -- "$(dirname -- "$0")" || exit 1
export PORTFOLIO_COLOR="${PORTFOLIO_COLOR:-always}"
exec ./bin/ssh-portfolio
