/*-- FOR TESTING PURPOSES - should be removed --*/

REINDEX TABLE public.profiles; --DANGEROUS

alter table profiles
  disable row level security;

drop policy "Public profiles are viewable by everyone." on profiles;
drop policy "Users can insert their own profile." on profiles;
drop policy "Users can update own profile." on profiles;

/*-- --*/

--Setting time zone of postgres 
ALTER DATABASE postgres SET TIMEZONE TO 'Asia/Kolkata';


alter table profiles
  enable row level security;

create policy "Public profiles are viewable by everyone." on profiles
  for select using (true);

create policy "Users can insert their own profile." on profiles
  for insert with check ((select auth.uid()) = id);

create policy "Users can update own profile." on profiles
  for update using ((select auth.uid()) = id);


create or replace function public.handle_new_user()
returns trigger
set search_path = ''
as $$
begin
  insert into public.profiles (id, name, role_id, email)
  values (new.id, new.raw_user_meta_data->>'name', (new.raw_user_meta_data->>'role_id')::integer, new.email);
  return new;
end;
$$ language plpgsql security definer;

CREATE OR REPLACE TRIGGER on_auth_user_confirmed
AFTER UPDATE ON auth.users
FOR EACH ROW
WHEN (NEW.email_confirmed_at IS NOT NULL AND OLD.email_confirmed_at IS NULL)
EXECUTE FUNCTION public.handle_new_user();