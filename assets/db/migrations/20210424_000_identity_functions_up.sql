--
-- Name: account(text); Type: FUNCTION; Schema: public; Owner: postgres
--
CREATE FUNCTION public.account(id text) RETURNS text
    LANGUAGE sql IMMUTABLE
    AS $_$
    SELECT CASE
       WHEN split_part($1, ':', 1) = '' THEN NULL
      ELSE split_part($1, ':', 1)
    END
    $_$;

ALTER FUNCTION public.account(id text) OWNER TO postgres;

--
-- Name: kind(text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.kind(id text) RETURNS text
    LANGUAGE sql IMMUTABLE
    AS $_$
    SELECT CASE
       WHEN split_part($1, ':', 2) = '' THEN NULL
      ELSE split_part($1, ':', 2)
    END
    $_$;


ALTER FUNCTION public.kind(id text) OWNER TO postgres;
