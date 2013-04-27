package core

import (

)

const SIPSeparatorNames_SEMICOLON = ";"
const SIPSeparatorNames_COLON = ":"
const SIPSeparatorNames_COMMA = ","
const SIPSeparatorNames_SLASH = "/"
const SIPSeparatorNames_SP = " "
const SIPSeparatorNames_EQUALS = "="
const SIPSeparatorNames_STAR = "*"
const SIPSeparatorNames_NEWLINE = "\r\n"
const SIPSeparatorNames_RETURN = "\n"
const SIPSeparatorNames_LESS_THAN = "<"
const SIPSeparatorNames_GREATER_THAN = ">"
const SIPSeparatorNames_AT = "@"
const SIPSeparatorNames_DOT = "."
const SIPSeparatorNames_QUESTION = "?"
const SIPSeparatorNames_POUND = "#"
const SIPSeparatorNames_AND = "&"
const SIPSeparatorNames_LPAREN = "("
const SIPSeparatorNames_RPAREN = ")"
const SIPSeparatorNames_DOUBLE_QUOTE = "\""
const SIPSeparatorNames_QUOTE = "'" //\'
const SIPSeparatorNames_HT = "\t"
const SIPSeparatorNames_PERCENT = "%"

const SIPHeaderNames_MIN_EXPIRES = "Min-Expires";  //1
const SIPHeaderNames_ERROR_INFO = "Error-Info";  //2
const SIPHeaderNames_MIME_VERSION = "MIME-Version"; //3
const SIPHeaderNames_IN_REPLY_TO= "In-Reply-To"; //4
const SIPHeaderNames_ALLOW= "Allow"; //5
const SIPHeaderNames_CONTENT_LANGUAGE= "Content-Language"; //6
const SIPHeaderNames_CALL_INFO= "Call-Info"; //7
const SIPHeaderNames_CSEQ="CSeq"; //8
const SIPHeaderNames_ALERT_INFO="Alert-Info"; //9
const SIPHeaderNames_ACCEPT_ENCODING="Accept-Encoding"; //10
const SIPHeaderNames_ACCEPT= "Accept"; //11
const SIPHeaderNames_ACCEPT_LANGUAGE= "Accept-Language"; //12
const SIPHeaderNames_RECORD_ROUTE= "Record-Route"; //13
const SIPHeaderNames_TIMESTAMP="Timestamp"; //14
const SIPHeaderNames_TO="To"; //15
const SIPHeaderNames_VIA="Via"; //16
const SIPHeaderNames_FROM="From"; //17
const SIPHeaderNames_CALL_ID="Call-ID"; //18
const SIPHeaderNames_AUTHORIZATION="Authorization"; //19
const SIPHeaderNames_PROXY_AUTHENTICATE="Proxy-Authenticate"; //20
const SIPHeaderNames_SERVER="Server"; //21
const SIPHeaderNames_UNSUPPORTED="Unsupported"; //22
const SIPHeaderNames_RETRY_AFTER="Retry-After"; //23
const SIPHeaderNames_CONTENT_TYPE="Content-Type"; //24
const SIPHeaderNames_CONTENT_ENCODING="Content-Encoding"; //25
const SIPHeaderNames_CONTENT_LENGTH="Content-Length"; //26
const SIPHeaderNames_ROUTE="Route"; //27
const SIPHeaderNames_CONTACT="Contact"; //28
const SIPHeaderNames_WWW_AUTHENTICATE="WWW-Authenticate"; //29
const SIPHeaderNames_MAX_FORWARDS="Max-Forwards"; //30
const SIPHeaderNames_ORGANIZATION="Organization"; //31
const SIPHeaderNames_PROXY_AUTHORIZATION="Proxy-Authorization"; //32
const SIPHeaderNames_PROXY_REQUIRE="Proxy-Require"; //33
const SIPHeaderNames_REQUIRE="Require"; //34
const SIPHeaderNames_CONTENT_DISPOSITION="Content-Disposition";//35
const SIPHeaderNames_SUBJECT="Subject";//36
const SIPHeaderNames_USER_AGENT="User-Agent";//37
const SIPHeaderNames_WARNING="Warning"; //38
const SIPHeaderNames_PRIORITY="Priority"; //39
const SIPHeaderNames_DATE="Date"; //40
const SIPHeaderNames_EXPIRES="Expires"; //41
const SIPHeaderNames_SUPPORTED = "Supported";//42
const SIPHeaderNames_AUTHENTICATION_INFO="Authentication-Info";//43
const SIPHeaderNames_REPLY_TO = "Reply-To"; //44
const SIPHeaderNames_RACK	= "RAck";//45
const SIPHeaderNames_RSEQ	= "RSeq";//46
const SIPHeaderNames_REASON = "Reason";//47
const SIPHeaderNames_SUBSCRIPTION_STATE = "Subscription-State";//48
const SIPHeaderNames_EVENT = "Event";//44
const SIPHeaderNames_ALLOW_EVENTS= "Allow-Events";//45
const SIPHeaderNames_REFER_TO = "Refer-To"//46
const SIPHeaderNames_K  = "K";
const SIPHeaderNames_C  = "C";
const SIPHeaderNames_E  = "E";
const SIPHeaderNames_F  = "F";
const SIPHeaderNames_I  = "I";
const SIPHeaderNames_M  = "M";
const SIPHeaderNames_L  = "L";
const SIPHeaderNames_S  = "S";
const SIPHeaderNames_T  = "T";
const SIPHeaderNames_V  = "V";
const SIPHeaderNames_R  = "R";

const SIPMethodNames_INVITE   = "INVITE";
const SIPMethodNames_ACK      = "ACK";
const SIPMethodNames_BYE      = "BYE";
const SIPMethodNames_SUBSCRIBE= "SUBSCRIBE";
const SIPMethodNames_NOTIFY   = "NOTIFY";
const SIPMethodNames_OPTIONS  = "OPTIONS";
const SIPMethodNames_REGISTER = "REGISTER";
const SIPMethodNames_MESSAGE  = "MESSAGE";

const SIPTransportNames_UDP = "udp";
const SIPTransportNames_TCP = "tcp";
const SIPTransportNames_TRANSPORT = "transport";
const SIPTransportNames_METHOD = "method";
const SIPTransportNames_USER = "user";
const SIPTransportNames_PHONE = "phone";
const SIPTransportNames_MADDR = "maddr";
const SIPTransportNames_TTL = "ttl";
const SIPTransportNames_LR = "lr";
const SIPTransportNames_SIP = "sip";
const SIPTransportNames_SIPS = "sips"; 
const SIPTransportNames_TEL = "tel";
const SIPTransportNames_POSTDIAL  =  "postdial";
const SIPTransportNames_PHONE_CONTEXT_TAG  = "context-tag";
const SIPTransportNames_ISUB 	  = "isub";
const SIPTransportNames_PROVIDER_TAG    = "provider-tag";

const SIPDateNames_GMT="GMT";
const SIPDateNames_MON="Mon";
const SIPDateNames_TUE="Tue";
const SIPDateNames_WED="Wed";
const SIPDateNames_THU="Thu";
const SIPDateNames_FRI="Fri";
const SIPDateNames_SAT="Sat";
const SIPDateNames_SUN="Sun";
const SIPDateNames_JAN="Jan";
const SIPDateNames_FEB="Feb";
const SIPDateNames_MAR="Mar";
const SIPDateNames_APR="Apr";
const SIPDateNames_MAY="May";
const SIPDateNames_JUN="Jun";
const SIPDateNames_JUL="Jul";
const SIPDateNames_AUG="Aug";
const SIPDateNames_SEP="Sep";
const SIPDateNames_OCT="Oct";
const SIPDateNames_NOV="Nov";
const SIPDateNames_DEC="Dec";

const SIPCalendar_AM	= 0
const SIPCalendar_AM_PM	= 9
const SIPCalendar_APRIL	= 3
const SIPCalendar_AUGUST= 7
const SIPCalendar_DATE	= 5
const SIPCalendar_DAY_OF_MONTH	= 5
const SIPCalendar_DAY_OF_WEEK	= 7
const SIPCalendar_DAY_OF_WEEK_IN_MONTH	= 8
const SIPCalendar_DAY_OF_YEAR	= 6
const SIPCalendar_DECEMBER	= 11
const SIPCalendar_DST_OFFSET	= 16
const SIPCalendar_ERA	= 0
const SIPCalendar_FEBRUARY	= 1
const SIPCalendar_FIELD_COUNT	= 17
const SIPCalendar_FRIDAY	= 6
const SIPCalendar_HOUR	= 10
const SIPCalendar_HOUR_OF_DAY	= 11
const SIPCalendar_JANUARY	= 0
const SIPCalendar_JULY	= 6
const SIPCalendar_JUNE	= 5
const SIPCalendar_MARCH	= 2
const SIPCalendar_MAY	= 4
const SIPCalendar_MILLISECOND	= 14
const SIPCalendar_MINUTE	= 12
const SIPCalendar_MONDAY	= 2
const SIPCalendar_MONTH		= 2
const SIPCalendar_NOVEMBER	= 10
const SIPCalendar_OCTOBER	= 9
const SIPCalendar_PM		= 1
const SIPCalendar_SATURDAY	= 7
const SIPCalendar_SECOND	= 13
const SIPCalendar_SEPTEMBER	= 8
const SIPCalendar_SUNDAY	= 1
const SIPCalendar_THURSDAY	= 5
const SIPCalendar_TUESDAY	= 3
const SIPCalendar_UNDECIMBER	= 12
const SIPCalendar_WEDNESDAY		= 4
const SIPCalendar_WEEK_OF_MONTH	= 4
const SIPCalendar_WEEK_OF_YEAR	= 3
const SIPCalendar_YEAR			= 1
const SIPCalendar_ZONE_OFFSET	= 15


	// Issue reported by larryb
    const SIPParameters_NEXT_NONCE = "nextnonce";
	const SIPParameters_TAG="tag";
	const SIPParameters_USERNAME="username";
	const SIPParameters_URI="uri";
	const SIPParameters_DOMAIN="domain";
	const SIPParameters_CNONCE="cnonce";
	const SIPParameters_PASSWORD="password";
	const SIPParameters_RESPONSE ="response";
	const SIPParameters_RESPONSE_AUTH="rspauth";
	const SIPParameters_OPAQUE="opaque";
	const SIPParameters_ALGORITHM="algorithm";
	const SIPParameters_DIGEST="Digest";
	const SIPParameters_SIGNED_BY="signed-by";
	const SIPParameters_SIGNATURE="signature";
	const SIPParameters_NONCE="nonce";
	// Issue reported by larryb
	const SIPParameters_NONCE_COUNT="nc";
	const SIPParameters_PUBKEY="pubkey";
	const SIPParameters_COOKIE="cookie";
	const SIPParameters_REALM="realm";
	const SIPParameters_VERSION="version";
	const SIPParameters_STALE="stale";
	const SIPParameters_QOP="qop";
	const SIPParameters_NC="nc";
	const SIPParameters_PURPOSE="purpose";                           
	const SIPParameters_CARD="card";                           
	const SIPParameters_INFO="info";                           
        const SIPParameters_ACTION="action";
        const SIPParameters_PROXY="proxy";
        const SIPParameters_REDIRECT="redirect"; 
	const SIPParameters_EXPIRES = "expires";
        const SIPParameters_Q="q";
	const SIPParameters_RENDER = "render";
	const SIPParameters_SESSION = "session";
	const SIPParameters_ICON	= "icon";
	const SIPParameters_ALERT   = "alert";
	const SIPParameters_HANDLING = "handling"; 
	const SIPParameters_REQUIRED = "required";
	const SIPParameters_OPTIONAL = "optional";
	const SIPParameters_EMERGENCY="emergency";
	const SIPParameters_URGENT="urgent";
	const SIPParameters_NORMAL="normal";
	const SIPParameters_NON_URGENT="non-urgent";
        const SIPParameters_DURATION="duration";
        const SIPParameters_BRANCH="branch";
        const SIPParameters_HIDDEN="hidden";
        const SIPParameters_RECEIVED="received";
        const SIPParameters_MADDR="maddr";
        const SIPParameters_TTL="ttl";
        const SIPParameters_TRANSPORT="transport";
        const SIPParameters_TEXT="text";
        const SIPParameters_CAUSE="cause";
        const SIPParameters_ID="id";

